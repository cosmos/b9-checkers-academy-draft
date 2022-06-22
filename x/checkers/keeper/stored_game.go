package keeper

import (
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredGame set a specific storedGame in the store from its index
func (k Keeper) SetStoredGame(ctx sdk.Context, storedGame types.StoredGame) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKey))
	b := k.cdc.MustMarshalBinaryBare(&storedGame)
	store.Set(types.KeyPrefix(storedGame.Index), b)
}

// GetStoredGame returns a storedGame from its index
func (k Keeper) GetStoredGame(ctx sdk.Context, index string) (val types.StoredGame, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveStoredGame removes a storedGame from the store
func (k Keeper) RemoveStoredGame(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllStoredGame returns all storedGame
func (k Keeper) GetAllStoredGame(ctx sdk.Context) (list []types.StoredGame) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredGame
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
