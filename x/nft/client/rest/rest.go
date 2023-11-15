package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// RegisterHandlers registers the NFT REST routes.
func RegisterHandlers(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

const (
	RestParamDenomID     = "denom-id"
	RestParamDenomName   = "denom-name"
	RestParamDenomSymbol = "denom-symbol"
	RestParamTokenID     = "token-id"
	RestParamOwner       = "owner"
)

type issueDenomReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Schema      string       `json:"schema"`
	Symbol      string       `json:"symbol"`
	Traits      string       `json:"traits"`
	Minter      string       `json:"minter"`
	Description string       `json:"description"`
	Data        string       `json:"data"`
}

type mintNFTReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Recipient string       `json:"recipient"`
	DenomID   string       `json:"denom_id"`
	Name      string       `json:"name"`
	URI       string       `json:"uri"`
	Data      string       `json:"data"`
}

type editNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name    string       `json:"name"`
	URI     string       `json:"uri"`
	Data    string       `json:"data"`
}

type transferNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	From    string       `json:"from"`
	To      string       `json:"to"`
}

type transferDenomReq struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Recipient string       `json:"recipient"`
	DenomID   string       `json:"denom_id"`
}

type approveNFTReq struct {
	DenomID          string       `json:"denom_id"`
	TokenID          string       `json:"token_id"`
	AddressToApprove string       `json:"address_to_approve"`
	BaseReq          rest.BaseReq `json:"base_req"`
}

type revokeNFTReq struct {
	BaseReq         rest.BaseReq `json:"base_req"`
	AddressToRevoke string       `json:"address_to_revoke"`
	DenomID         string       `json:"denom_id"`
	TokenID         string       `json:"token_id"`
}

type burnNFTReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	DnomId  string       `json:"denom_id"`
	TokenID string       `json:"token_id"`
}

type approveAllRequest struct {
	BaseReq          rest.BaseReq `json:"base_req"`
	ApprovedOperator string       `json:"approved_operator"`
	Approved         bool         `json:"approved"`
}

type queryIsApprovedForAllRequest struct {
	Owner    string `json:"owner"`
	Operator string `json:"operator"`
}

type queryDenomsRequest struct {
	Pagination query.PageRequest `json:"pagination"`
}

type queryCollectionRequest struct {
	DenomID    string            `json:"denom_id"`
	Pagination query.PageRequest `json:"pagination"`
}

type queryCollectionsByDenomIDsRequest struct {
	DenomIDs []string `json:"denom_ids"`
}

type queryOwnerRequest struct {
	DenomID      string            `json:"denom_id,omitempty"`
	OwnerAddress string            `json:"owner_address"`
	Pagination   query.PageRequest `json:"pagination"`
}
