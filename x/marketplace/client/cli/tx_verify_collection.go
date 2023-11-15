package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/CudoVentures/cudos-node/x/marketplace/types"
)

var _ = strconv.Itoa(0)

func CmdVerifyCollection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-collection [id]",
		Short: "Verify collection",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			collectionID, err := strconv.ParseUint(argId, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgVerifyCollection(
				clientCtx.GetFromAddress().String(),
				collectionID,
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
