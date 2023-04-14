package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ CheckersHooks = MultiCheckersHooks{}

type MultiCheckersHooks []CheckersHooks

func NewMultiCheckersHooks(hooks ...CheckersHooks) MultiCheckersHooks {
	return hooks
}

func (h MultiCheckersHooks) AfterPlayerInfoChanged(ctx sdk.Context, playerInfo PlayerInfo) {
	for i := range h {
		h[i].AfterPlayerInfoChanged(ctx, playerInfo)
	}
}
