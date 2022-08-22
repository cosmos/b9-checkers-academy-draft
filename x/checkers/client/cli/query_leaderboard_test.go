package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/b9lab/checkers/testutil/network"
	"github.com/b9lab/checkers/testutil/nullify"
	"github.com/b9lab/checkers/x/checkers/client/cli"
	"github.com/b9lab/checkers/x/checkers/types"
)

func networkWithLeaderboardObjects(t *testing.T) (*network.Network, types.Leaderboard) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	leaderboard := types.Leaderboard{}
	nullify.Fill(&leaderboard)
	state.Leaderboard = leaderboard
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.Leaderboard
}

func TestShowLeaderboard(t *testing.T) {
	net, obj := networkWithLeaderboardObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.Leaderboard
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowLeaderboard(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetLeaderboardResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Leaderboard)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Leaderboard),
				)
			}
		})
	}
}
