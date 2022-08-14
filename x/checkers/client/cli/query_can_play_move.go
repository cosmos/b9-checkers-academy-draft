package cli

import (
	"strconv"

	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCanPlayMove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "can-play-move [game-index] [player] [from-x] [from-y] [to-x] [to-y]",
		Short: "Query canPlayMove",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGameIndex := args[0]
			reqPlayer := args[1]
			reqFromX, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			reqFromY, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			reqToX, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			reqToY, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCanPlayMoveRequest{

				GameIndex: reqGameIndex,
				Player:    reqPlayer,
				FromX:     reqFromX,
				FromY:     reqFromY,
				ToX:       reqToX,
				ToY:       reqToY,
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
