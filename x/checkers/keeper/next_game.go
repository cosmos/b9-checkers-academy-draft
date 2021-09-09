package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

// SetNextGame set nextGame in the store
func (k Keeper) SetNextGame(ctx sdk.Context, nextGame types.NextGame) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextGameKey))
	b := k.cdc.MustMarshalBinaryBare(&nextGame)
	store.Set([]byte{0}, b)
}

// GetNextGame returns nextGame
func (k Keeper) GetNextGame(ctx sdk.Context) (val types.NextGame, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextGameKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveNextGame removes nextGame from the store
func (k Keeper) RemoveNextGame(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextGameKey))
	store.Delete([]byte{0})
}
