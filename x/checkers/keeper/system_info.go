package keeper

import (
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSystemInfo set systemInfo in the store
func (k Keeper) SetSystemInfo(ctx sdk.Context, systemInfo types.SystemInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKey))
	b := k.cdc.MustMarshal(&systemInfo)
	store.Set([]byte{0}, b)
}

// GetSystemInfo returns systemInfo
func (k Keeper) GetSystemInfo(ctx sdk.Context) (val types.SystemInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSystemInfo removes systemInfo from the store
func (k Keeper) RemoveSystemInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKey))
	store.Delete([]byte{0})
}
