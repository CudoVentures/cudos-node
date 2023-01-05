package keeper

import (
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/CudoVentures/cudos-node/x/marketplace/types"
	"github.com/CudoVentures/cudos-node/x/nft/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
		nftKeeper  types.NftKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper, nftKeeper types.NftKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper, nftKeeper: nftKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) isCollectionPublished(ctx sdk.Context, denomID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyCollectionDenomID(denomID))
}

func (k Keeper) PublishCollection(ctx sdk.Context, collection types.Collection) (uint64, error) {
	denom, err := k.nftKeeper.GetDenom(ctx, collection.DenomId)
	if err != nil {
		return 0, err
	}

	if denom.Creator != collection.Owner {
		return 0, sdkerrors.Wrapf(types.ErrNotDenomOwner, "Owner of denom %s is %s", collection.DenomId, denom.Creator)
	}

	if k.isCollectionPublished(ctx, collection.DenomId) {
		return 0, sdkerrors.Wrapf(types.ErrCollectionAlreadyPublished, "Collection for denom %s is already published", collection.DenomId)
	}

	collectionID := k.AppendCollection(ctx, collection)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyCollectionDenomID(collection.DenomId), types.Uint64ToBytes(collectionID))

	return collectionID, nil
}

func (k Keeper) isApprovedNftAddress(nftVal exported.NFT, owner string) bool {
	approvedAddresses := nftVal.GetApprovedAddresses()
	for _, addr := range approvedAddresses {
		if addr == owner {
			return true
		}
	}
	return false
}

func (k Keeper) PublishNFT(ctx sdk.Context, nft types.Nft) (uint64, error) {
	if _, err := k.nftKeeper.GetDenom(ctx, nft.DenomId); err != nil {
		return 0, err
	}

	nftVal, err := k.nftKeeper.GetNFT(ctx, nft.DenomId, nft.TokenId)
	if err != nil {
		return 0, err
	}

	publisher, err := sdk.AccAddressFromBech32(nft.Owner)
	if err != nil {
		return 0, err
	}

	if nftVal.GetOwner().String() == nft.Owner ||
		k.nftKeeper.IsApprovedOperator(ctx, nftVal.GetOwner(), publisher) ||
		k.isApprovedNftAddress(nftVal, nft.Owner) {

		store := ctx.KVStore(k.storeKey)
		key := types.KeyNftDenomTokenID(nft.DenomId, nft.TokenId)
		if b := store.Get(key); len(b) > 0 {
			return 0, sdkerrors.Wrapf(types.ErrNftAlreadyPublished, "nft with token id (%s) from denom (%s) already published for sale", nft.TokenId, nft.DenomId)
		}

		if err := k.nftKeeper.SoftLockNFT(ctx, types.ModuleName, nft.DenomId, nft.TokenId); err != nil {
			return 0, err
		}

		nftID := k.AppendNft(ctx, nft)

		store.Set(key, types.Uint64ToBytes(nftID))

		return nftID, nil
	}

	return 0, sdkerrors.Wrapf(types.ErrNotNftOwner, "%s not nft owner or approved operator for token id (%s) from denom (%s)", nft.Owner, nft.TokenId, nft.DenomId)
}

func (k Keeper) BuyNFT(ctx sdk.Context, nftID uint64, buyer sdk.AccAddress) (types.Nft, error) {
	nft, found := k.GetNft(ctx, nftID)
	if !found {
		return types.Nft{}, sdkerrors.Wrapf(types.ErrNftNotFound, "nft with id (%d) is not found for sale", nftID)
	}

	if nft.Owner == buyer.String() {
		return types.Nft{}, sdkerrors.Wrap(types.ErrCannotBuyOwnNft, "cannot buy own nft")
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, sdk.NewCoins(nft.Price)); err != nil {
		return types.Nft{}, err
	}

	seller, err := sdk.AccAddressFromBech32(nft.Owner)
	if err != nil {
		return types.Nft{}, err
	}

	if err := k.handleSale(ctx, nft.DenomId, nft.TokenId, buyer, seller, nft.Price); err != nil {
		return types.Nft{}, err
	}

	k.RemoveNft(ctx, nftID)
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNftDenomTokenID(nft.DenomId, nft.TokenId))
	return nft, nil
}

func (k Keeper) MintNFT(ctx sdk.Context, denomID, name, uri, data string, price sdk.Coin, recipient sdk.AccAddress, sender sdk.AccAddress) (string, error) {
	denom, err := k.nftKeeper.GetDenom(ctx, denomID)
	if err != nil {
		return "", err
	}

	collection, found := k.GetCollectionByDenomID(ctx, denomID)
	if !found {
		return "", sdkerrors.Wrapf(types.ErrCollectionNotFound, "collection %s not published for sale", denomID)
	}

	if !collection.Verified {
		return "", sdkerrors.Wrapf(types.ErrCollectionIsUnverified, "collection %d is not verified", collection.Id)
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(price)); err != nil {
		return "", err
	}

	if err := k.DistributeRoyalties(ctx, price, denom.Creator, collection.MintRoyalties); err != nil {
		return "", err
	}

	return k.nftKeeper.MintNFT(ctx, denomID, name, uri, data, sender, recipient)
}

func (k Keeper) RemoveNFT(ctx sdk.Context, nftID uint64, owner sdk.AccAddress) (types.Nft, error) {
	nft, found := k.GetNft(ctx, nftID)
	if !found {
		return types.Nft{}, sdkerrors.Wrapf(types.ErrNftNotFound, "nft with id (%d) is not found for sale", nftID)
	}

	if nft.Owner != owner.String() {
		return types.Nft{}, sdkerrors.Wrapf(types.ErrNotNftOwner, "not owner of (%d)", nftID)
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNftDenomTokenID(nft.DenomId, nft.TokenId))

	k.RemoveNft(ctx, nftID)

	if err := k.nftKeeper.SoftUnlockNFT(ctx, types.ModuleName, nft.DenomId, nft.TokenId); err != nil {
		return types.Nft{}, err
	}

	return nft, nil
}

func (k Keeper) CreateCollection(ctx sdk.Context, sender sdk.AccAddress, id, name, schema, symbol, traits, description, minter, data string, mintRoyalties, resaleRoyalties []types.Royalty, verified bool) (uint64, error) {
	if err := k.nftKeeper.IssueDenom(ctx, id, name, schema, symbol, traits, minter, description, data, sender); err != nil {
		return 0, err
	}

	return k.PublishCollection(ctx, types.NewCollection(id, mintRoyalties, resaleRoyalties, sender.String(), verified))
}

func (k Keeper) GetCollectionByDenomID(ctx sdk.Context, denomID string) (types.Collection, bool) {
	store := ctx.KVStore(k.storeKey)
	collectionIDBytes := store.Get(types.KeyCollectionDenomID(denomID))
	if collectionIDBytes == nil {
		return types.Collection{}, false
	}
	return k.GetCollection(ctx, types.BytesToUint64(collectionIDBytes))
}

func getProportion(totalCoin sdk.Coin, ratio sdk.Dec) sdk.Coin {
	return sdk.NewCoin(totalCoin.Denom, totalCoin.Amount.ToDec().Mul(ratio).Quo(sdk.NewDec(100)).TruncateInt())
}

func (k Keeper) DistributeRoyalties(ctx sdk.Context, price sdk.Coin, seller string, royalties []types.Royalty) error {
	if len(royalties) == 0 {
		return nil
	}

	amountLeft := price.Amount

	for _, royalty := range royalties {

		royaltyReceiver, err := sdk.AccAddressFromBech32(royalty.Address)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid royalty address (%s): %s", royalty.Address, err)
		}

		portion := getProportion(price, royalty.Percent)
		amountLeft = amountLeft.Sub(portion.Amount)

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, royaltyReceiver, sdk.NewCoins(portion)); err != nil {
			return err
		}
	}

	if amountLeft.GT(sdk.NewInt(0)) {
		sellerAddr, err := sdk.AccAddressFromBech32(seller)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid seller address (%s): %s", seller, err)
		}

		return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sellerAddr, sdk.NewCoins(sdk.NewCoin(price.Denom, amountLeft)))
	}

	return nil
}

func (k Keeper) GetCollectionStatus(ctx sdk.Context, id uint64) (bool, error) {
	collection, found := k.GetCollection(ctx, id)
	if !found {
		return false, sdkerrors.Wrapf(types.ErrCollectionNotFound, "collection with id %d not found", id)
	}
	return collection.Verified, nil
}

func (k Keeper) SetCollectionStatus(ctx sdk.Context, id uint64, verified bool) error {
	collection, found := k.GetCollection(ctx, id)
	if !found {
		return sdkerrors.Wrapf(types.ErrCollectionNotFound, "collection with id %d not found", id)
	}
	collection.Verified = verified
	k.SetCollection(ctx, collection)
	return nil
}

func (k Keeper) SetCollectionRoyalties(ctx sdk.Context, sender string, id uint64, mintRoyalties, resaleRoyalties []types.Royalty) error {
	collection, found := k.GetCollection(ctx, id)
	if !found {
		return sdkerrors.Wrapf(types.ErrCollectionNotFound, "collection with id %d not found", id)
	}

	if collection.Owner != sender {
		return sdkerrors.Wrapf(types.ErrNotCollectionOwner, "owner of collection %d is %s, not %s", id, collection.Owner, sender)
	}

	collection.MintRoyalties = mintRoyalties
	collection.ResaleRoyalties = resaleRoyalties
	k.SetCollection(ctx, collection)
	return nil
}

func (k Keeper) SetNftPrice(ctx sdk.Context, sender string, id uint64, price sdk.Coin) (types.Nft, error) {
	nft, found := k.GetNft(ctx, id)
	if !found {
		return types.Nft{}, sdkerrors.Wrapf(types.ErrNftNotFound, "NFT with id %d not found", id)
	}

	if nft.Owner != sender {
		return types.Nft{}, sdkerrors.Wrapf(types.ErrNotCollectionOwner, "owner of NFT %d is %s, not %s", id, nft.Owner, sender)
	}

	nft.Price = price
	k.SetNft(ctx, nft)

	return nft, nil
}

func (k Keeper) GetAdmins(ctx sdk.Context) ([]string, error) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.KeyAdmins())
	if b == nil {
		return []string{}, nil
	}

	var admins types.Admins
	k.cdc.MustUnmarshal(b, &admins)

	return admins.Addresses, nil
}

func (k Keeper) IsAdmin(ctx sdk.Context, address string) error {
	admins, err := k.GetAdmins(ctx)
	if err != nil {
		return err
	}

	for _, admin := range admins {
		if admin == address {
			return nil
		}
	}

	return sdkerrors.Wrapf(types.ErrNotAdmin, "'%s' is not admin", address)
}

func (k Keeper) isCudosAdmin(ctx sdk.Context, address string) error {
	accAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	balance := k.bankKeeper.GetBalance(ctx, accAddr, types.AdminDenom)
	if balance.IsPositive() {
		return nil
	}

	return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Insufficient permissions. Address '%s' has no %s tokens", address, types.AdminDenom)
}

func (k Keeper) setAdmins(ctx sdk.Context, admins []string) {
	b := k.cdc.MustMarshal(&types.Admins{Addresses: admins})
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyAdmins(), b)
}

func (k Keeper) AddAdmin(ctx sdk.Context, admin, creator string) error {
	if err := k.isCudosAdmin(ctx, creator); err != nil {
		return err
	}

	admins, err := k.GetAdmins(ctx)
	if err != nil {
		return err
	}

	for _, address := range admins {
		if address == admin {
			return sdkerrors.Wrapf(types.ErrAlreadyAdmin, "'%s' is already admin.", admin)
		}
	}

	admins = append(admins, admin)

	k.setAdmins(ctx, admins)

	return nil
}

func (k Keeper) RemoveAdmin(ctx sdk.Context, admin, creator string) error {
	if err := k.isCudosAdmin(ctx, creator); err != nil {
		return err
	}

	admins, err := k.GetAdmins(ctx)
	if err != nil {
		return err
	}

	for i, address := range admins {
		if address == admin {
			k.setAdmins(ctx, append(admins[:i], admins[i+1:]...))
			return nil
		}
	}

	return sdkerrors.Wrapf(types.ErrNotAdmin, "'%s' is not admin", admin)
}

func (k Keeper) PublishAuction(ctx sdk.Context, a types.Auction) (uint64, error) {
	// check if nft is already listed
	store := ctx.KVStore(k.storeKey)
	key := types.KeyNftDenomTokenID(a.DenomId, a.TokenId)
	if len(store.Get(key)) > 0 {
		return 0, sdkerrors.Wrapf(types.ErrNftAlreadyPublished, "nft is already published")
	}

	creator, err := sdk.AccAddressFromBech32(a.Creator)
	if err != nil {
		return 0, err
	}

	nft, err := k.nftKeeper.GetNFT(ctx, a.DenomId, a.TokenId)
	if err != nil {
		return 0, err
	}

	// check if user is authorized - must be nft owner or approved operator
	if nft.GetOwner().String() != a.Creator &&
		!k.nftKeeper.IsApprovedOperator(ctx, nft.GetOwner(), creator) &&
		!k.isApprovedNftAddress(nft, a.Creator) {
		return 0, sdkerrors.Wrapf(types.ErrNotNftOwner, "%s not nft owner or approved operator for token id (%s) from denom (%s)", a.Creator, a.TokenId, a.DenomId)
	}

	// lock nft during auction
	if err := k.nftKeeper.SoftLockNFT(ctx, types.ModuleName, a.DenomId, a.TokenId); err != nil {
		return 0, err
	}

	// save auction to store
	auctionID, err := k.AppendAuction(ctx, a)
	if err != nil {
		return 0, err
	}
	// mark nft as listed
	store.Set(key, types.Uint64ToBytes(auctionID))

	return auctionID, nil
}

func (k Keeper) PlaceBid(ctx sdk.Context, auctionID uint64, bid types.Bid) error {
	a, found := k.GetAuction(ctx, auctionID)
	if !found {
		return sdkerrors.Wrapf(types.ErrAuctionNotFound, "auction with id (%d) does not exist", auctionID)
	}

	if a.Creator == bid.Bidder {
		return sdkerrors.Wrap(types.ErrCannotBuyOwnNft, "cannot bid own auctions")
	}

	if ctx.BlockTime().After(a.EndTime) {
		return sdkerrors.Wrapf(types.ErrAuctionExpired, "cannot place a bid for inactive auction %d", a.Id)
	}

	auctionType, err := a.GetAuctionType()
	if err != nil {
		return err
	}

	switch at := auctionType.(type) {
	case *types.EnglishAuction:
		if err := k.handleBidEnglishAuction(ctx, at, bid); err != nil {
			return err
		}
		a.SetAuctionType(at)
	}

	return k.SetAuction(ctx, a)
}

func (k Keeper) handleBidEnglishAuction(ctx sdk.Context, a *types.EnglishAuction, bid types.Bid) error {
	if bid.Amount.IsLT(a.MinPrice) {
		return sdkerrors.Wrap(types.ErrInvalidPrice, "bid is lower than auction minimum price")
	}

	if a.CurrentBid != nil {
		if a.CurrentBid.Amount.IsGTE(bid.Amount) {
			return sdkerrors.Wrap(types.ErrInvalidPrice, "bid is lower than current bid")
		}

		bidder, err := sdk.AccAddressFromBech32(a.CurrentBid.Bidder)
		if err != nil {
			return err
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.NewCoins(a.CurrentBid.Amount)); err != nil {
			return err
		}
	}

	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	if err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, sdk.NewCoins(bid.Amount)); err != nil {
		return err
	}

	a.CurrentBid = &bid

	return nil
}

func (k Keeper) AuctionEndBlocker(ctx sdk.Context) error {
	for _, a := range k.GetAllAuction(ctx) {
		auctionType, err := a.GetAuctionType()
		if err != nil {
			return err
		}

		switch at := auctionType.(type) {
		case *types.EnglishAuction:
			if err := k.handleEndBlockEnglishAuction(ctx, a, at); err != nil {
				return err
			}
		}
	}

	// TODO dutch auctions
	// 1. on PublishAuction calculate how many blocks until expiration
	// 2. we need to lower the price every block by: (startPrice - endPrice) / blocks
	// 3. keep that number (or calculate live and have 3rd field CurrentPrice)
	return nil
}

func (k Keeper) handleEndBlockEnglishAuction(ctx sdk.Context, a types.Auction, at *types.EnglishAuction) error {
	if ctx.BlockTime().Before(a.EndTime) {
		return fmt.Errorf("auction has not expired yet")
	}

	events := sdk.Events{
		sdk.NewEvent(
			types.EventAuctionExpiredType,
			sdk.NewAttribute(types.AttributeAuctionID, strconv.FormatUint(a.Id, 10)),
		),
	}

	if at.CurrentBid == nil {
		if err := k.nftKeeper.SoftUnlockNFT(ctx, types.ModuleName, a.DenomId, a.TokenId); err != nil {
			return err
		}
	} else {
		buyer, err := sdk.AccAddressFromBech32(at.CurrentBid.Bidder)
		if err != nil {
			return err
		}

		seller, err := sdk.AccAddressFromBech32(a.Creator)
		if err != nil {
			return err
		}

		if err := k.handleSale(ctx, a.DenomId, a.TokenId, buyer, seller, at.CurrentBid.Amount); err != nil {
			return err
		}

		events.AppendEvent(sdk.NewEvent(
			types.EventBuyNftType,
			sdk.NewAttribute(types.AttributeKeyDenomID, a.DenomId),
			sdk.NewAttribute(types.AttributeKeyTokenID, a.TokenId),
			sdk.NewAttribute(types.AttributeAuctionID, strconv.FormatUint(a.Id, 10)),
			sdk.NewAttribute(types.AttributeKeyBuyer, at.CurrentBid.Bidder),
		))

	}

	k.RemoveAuction(ctx, a.Id)
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyNftDenomTokenID(a.DenomId, a.TokenId))

	ctx.EventManager().EmitEvents(events)
	return nil
}

func (k Keeper) handleSale(ctx sdk.Context, denomId string, tokenId string, buyer sdk.AccAddress, seller sdk.AccAddress, amount sdk.Coin) error {
	collection, found := k.GetCollectionByDenomID(ctx, denomId)
	if !found || len(collection.ResaleRoyalties) == 0 {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, seller, sdk.NewCoins(amount)); err != nil {
			return err
		}
	}

	if err := k.DistributeRoyalties(ctx, amount, seller.String(), collection.ResaleRoyalties); err != nil {
		return err
	}

	baseNft, err := k.nftKeeper.GetBaseNFT(ctx, denomId, tokenId)
	if err != nil {
		return err
	}

	if err := k.nftKeeper.SoftUnlockNFT(ctx, types.ModuleName, denomId, tokenId); err != nil {
		return err
	}

	k.nftKeeper.TransferNftInternal(ctx, denomId, tokenId, seller, buyer, baseNft)
	return nil
}
