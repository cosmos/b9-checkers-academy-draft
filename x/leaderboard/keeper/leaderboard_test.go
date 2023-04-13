package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/testutil/nullify"
	"github.com/b9lab/checkers/x/leaderboard/keeper"
	"github.com/b9lab/checkers/x/leaderboard/types"
)

func createTestLeaderboard(keeper *keeper.Keeper, ctx sdk.Context) types.Leaderboard {
	item := types.Leaderboard{}
	keeper.SetLeaderboard(ctx, item)
	return item
}

func TestLeaderboardGet(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	item := createTestLeaderboard(keeper, ctx)
	rst := keeper.GetLeaderboard(ctx)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLeaderboardRemove(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	createTestLeaderboard(keeper, ctx)
	keeper.RemoveLeaderboard(ctx)
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, r, "Leaderboard not found")
	}()
	keeper.GetLeaderboard(ctx)
}
