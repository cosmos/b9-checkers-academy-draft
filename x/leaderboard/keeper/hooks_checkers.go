package keeper

import (
	"fmt"

	checkerstypes "github.com/b9lab/checkers/x/checkers/types"
	"github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ checkerstypes.CheckersHooks = Hooks{}

func (h Hooks) AfterPlayerInfoChanged(ctx sdk.Context, playerInfo checkerstypes.PlayerInfo) {
	candidate, err := types.MakeCandidateFromPlayerInfo(playerInfo)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	if candidate.WonCount < 1 {
		return
	}
	h.k.SetCandidate(ctx, candidate)
}
