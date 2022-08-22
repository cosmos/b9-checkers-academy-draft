package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/testutil/nullify"
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
)

func createTestLeaderboard(keeper *keeper.Keeper, ctx sdk.Context) types.Leaderboard {
	item := types.Leaderboard{}
	keeper.SetLeaderboard(ctx, item)
	return item
}

func TestLeaderboardGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	item := createTestLeaderboard(keeper, ctx)
	rst, found := keeper.GetLeaderboard(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLeaderboardRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	createTestLeaderboard(keeper, ctx)
	keeper.RemoveLeaderboard(ctx)
	_, found := keeper.GetLeaderboard(ctx)
	require.False(t, found)
}
