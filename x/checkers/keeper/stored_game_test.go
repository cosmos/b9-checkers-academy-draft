package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func createNStoredGame(keeper *Keeper, ctx sdk.Context, n int) []types.StoredGame {
	items := make([]types.StoredGame, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetStoredGame(ctx, items[i])
	}
	return items
}

func TestStoredGameGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredGame(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestStoredGameRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredGame(ctx, item.Index)
		_, found := keeper.GetStoredGame(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestStoredGameGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllStoredGame(ctx))
}
