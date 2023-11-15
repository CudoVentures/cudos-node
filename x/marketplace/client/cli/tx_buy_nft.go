package cli

import (
	"strconv"

	"github.com/CudoVentures/cudos-node/x/marketplace/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdBuyNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-nft [id]",
		Short: "Buy NFT",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nftID, err := strconv.ParseUint(argId, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyNft(
				clientCtx.GetFromAddress().String(),
				nftID,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
