package keeper

import (
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPlayerInfo set a specific playerInfo in the store from its index
func (k Keeper) SetPlayerInfo(ctx sdk.Context, playerInfo types.PlayerInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKeyPrefix))
	b := k.cdc.MustMarshal(&playerInfo)
	store.Set(types.PlayerInfoKey(
		playerInfo.Index,
	), b)
}

// GetPlayerInfo returns a playerInfo from its index
func (k Keeper) GetPlayerInfo(
	ctx sdk.Context,
	index string,

) (val types.PlayerInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKeyPrefix))

	b := store.Get(types.PlayerInfoKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePlayerInfo removes a playerInfo from the store
func (k Keeper) RemovePlayerInfo(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKeyPrefix))
	store.Delete(types.PlayerInfoKey(
		index,
	))
}

// GetAllPlayerInfo returns all playerInfo
func (k Keeper) GetAllPlayerInfo(ctx sdk.Context) (list []types.PlayerInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PlayerInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PlayerInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
