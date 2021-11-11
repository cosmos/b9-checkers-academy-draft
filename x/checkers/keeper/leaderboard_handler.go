package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func (k *Keeper) MustAddToLeaderboard(ctx sdk.Context, winnerInfo types.PlayerInfo) types.Leaderboard {
	leaderboard, found := k.GetLeaderboard(ctx)
	if !found {
		panic("Leaderboard not found")
	}
	err := leaderboard.AddCandidateAndSort(ctx, winnerInfo)
	if err != nil {
		panic(fmt.Sprintf(types.ErrCannotAddToLeaderboard.Error(), err.Error()))
	}
	k.SetLeaderboard(ctx, leaderboard)
	return leaderboard
}
