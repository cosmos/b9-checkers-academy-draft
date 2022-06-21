package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdRejectGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reject-game [idValue]",
		Short: "Broadcast message rejectGame",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsIdValue := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRejectGame(clientCtx.GetFromAddress().String(), string(argsIdValue))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
