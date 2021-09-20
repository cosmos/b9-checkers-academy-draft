package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

var _ = strconv.Itoa(0)

func CmdCanPlayMove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "can-play-move [idValue] [player] [fromX] [fromY] [toX] [toY]",
		Short: "Query canPlayMove",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			reqIdValue := string(args[0])
			reqPlayer := string(args[1])
			reqFromX, _ := strconv.ParseUint(args[2], 10, 64)
			reqFromY, _ := strconv.ParseUint(args[3], 10, 64)
			reqToX, _ := strconv.ParseUint(args[4], 10, 64)
			reqToY, _ := strconv.ParseUint(args[5], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCanPlayMoveRequest{

				IdValue: string(reqIdValue),
				Player:  string(reqPlayer),
				FromX:   uint64(reqFromX),
				FromY:   uint64(reqFromY),
				ToX:     uint64(reqToX),
				ToY:     uint64(reqToY),
			}

			res, err := queryClient.CanPlayMove(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
