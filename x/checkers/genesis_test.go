package checkers_test

import (
	"testing"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/testutil/nullify"
	"github.com/b9lab/checkers/x/checkers"
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 27,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PlayerInfoList: []types.PlayerInfo{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		Leaderboard: types.Leaderboard{
			Winners: []types.WinningPlayer{
				{
					PlayerAddress: "cosmos123",
				},
				{
					PlayerAddress: "cosmos456",
				},
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, genesisState)
	got := checkers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	require.ElementsMatch(t, genesisState.PlayerInfoList, got.PlayerInfoList)
	require.Equal(t, genesisState.Leaderboard, got.Leaderboard)
	// this line is used by starport scaffolding # genesis/test/assert
}
