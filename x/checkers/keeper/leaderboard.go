package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetLeaderboard set leaderboard in the store
func (k Keeper) SetLeaderboard(ctx sdk.Context, leaderboard types.Leaderboard) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	b := k.cdc.MustMarshalBinaryBare(&leaderboard)
	store.Set([]byte{0}, b)
}

// GetLeaderboard returns leaderboard
func (k Keeper) GetLeaderboard(ctx sdk.Context) (val types.Leaderboard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))

	b := store.Get([]byte{0})
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveLeaderboard removes leaderboard from the store
func (k Keeper) RemoveLeaderboard(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	store.Delete([]byte{0})
}
