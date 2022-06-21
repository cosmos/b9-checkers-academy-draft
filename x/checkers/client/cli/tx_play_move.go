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

func CmdPlayMove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play-move [idValue] [fromX] [fromY] [toX] [toY]",
		Short: "Broadcast message playMove",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsIdValue := string(args[0])
			argsFromX, _ := strconv.ParseUint(args[1], 10, 64)
			argsFromY, _ := strconv.ParseUint(args[2], 10, 64)
			argsToX, _ := strconv.ParseUint(args[3], 10, 64)
			argsToY, _ := strconv.ParseUint(args[4], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPlayMove(clientCtx.GetFromAddress().String(), string(argsIdValue), uint64(argsFromX), uint64(argsFromY), uint64(argsToX), uint64(argsToY))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
