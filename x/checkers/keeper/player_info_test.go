package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func createNPlayerInfo(keeper *Keeper, ctx sdk.Context, n int) []types.PlayerInfo {
	items := make([]types.PlayerInfo, n)
	for i := range items {
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetPlayerInfo(ctx, items[i])
	}
	return items
}

func TestPlayerInfoGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPlayerInfo(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestPlayerInfoRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePlayerInfo(ctx, item.Index)
		_, found := keeper.GetPlayerInfo(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestPlayerInfoGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPlayerInfo(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllPlayerInfo(ctx))
}
