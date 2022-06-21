package keeper

import (
	"fmt"

	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
