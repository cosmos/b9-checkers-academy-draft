package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func createTestLeaderboard(keeper *Keeper, ctx sdk.Context) types.Leaderboard {
	item := types.Leaderboard{}
	keeper.SetLeaderboard(ctx, item)
	return item
}

func TestLeaderboardGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	item := createTestLeaderboard(keeper, ctx)
	rst, found := keeper.GetLeaderboard(ctx)
	assert.True(t, found)
	assert.Equal(t, item, rst)
}
func TestLeaderboardRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	createTestLeaderboard(keeper, ctx)
	keeper.RemoveLeaderboard(ctx)
	_, found := keeper.GetLeaderboard(ctx)
	assert.False(t, found)
}
