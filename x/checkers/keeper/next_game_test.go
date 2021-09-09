package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func createTestNextGame(keeper *Keeper, ctx sdk.Context) types.NextGame {
	item := types.NextGame{
		Creator: "any",
	}
	keeper.SetNextGame(ctx, item)
	return item
}

func TestNextGameGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	item := createTestNextGame(keeper, ctx)
	rst, found := keeper.GetNextGame(ctx)
	assert.True(t, found)
	assert.Equal(t, item, rst)
}
func TestNextGameRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	createTestNextGame(keeper, ctx)
	keeper.RemoveNextGame(ctx)
	_, found := keeper.GetNextGame(ctx)
	assert.False(t, found)
}
