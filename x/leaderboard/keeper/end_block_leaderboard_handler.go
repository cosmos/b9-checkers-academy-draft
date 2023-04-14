package keeper

import (
	"github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CollectSortAndClipLeaderboard(ctx sdk.Context) {
	leaderboard := k.GetLeaderboard(ctx)
	updated := types.AddCandidatesAtNow(leaderboard.Winners, ctx.BlockTime(), k.GetAllCandidates(ctx))
	params := k.GetParams(ctx)
	if params.Length < uint64(len(updated)) {
		updated = updated[:params.Length]
	}
	leaderboard.Winners = updated
	k.SetLeaderboard(ctx, leaderboard)
}
