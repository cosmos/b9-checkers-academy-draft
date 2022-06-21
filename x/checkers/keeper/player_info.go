package keeper

import (
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPlayerInfo set a specific playerInfo in the store from its index
func (k Keeper) SetPlayerInfo(ctx sdk.Context, playerInfo types.PlayerInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKey))
	b := k.cdc.MustMarshalBinaryBare(&playerInfo)
	store.Set(types.KeyPrefix(playerInfo.Index), b)
}

// GetPlayerInfo returns a playerInfo from its index
func (k Keeper) GetPlayerInfo(ctx sdk.Context, index string) (val types.PlayerInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemovePlayerInfo removes a playerInfo from the store
func (k Keeper) RemovePlayerInfo(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllPlayerInfo returns all playerInfo
func (k Keeper) GetAllPlayerInfo(ctx sdk.Context) (list []types.PlayerInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PlayerInfo
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
