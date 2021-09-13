package keeper

import (
	"context"
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func (k msgServer) RejectGame(goCtx context.Context, msg *types.MsgRejectGame) (*types.MsgRejectGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.IdValue)
	if !found {
		return nil, errors.New("Game not found " + msg.IdValue)
	}
	fullGame := storedGame.ToFullGame()

	// Is it an expected player? And did the player already play?
	if strings.Compare(storedGame.Red, msg.Creator) == 0 {
		if 1 < fullGame.MoveCount {
			return nil, errors.New("Red player has already played, and cannot reject")
		}
	} else if strings.Compare(storedGame.Black, msg.Creator) == 0 {
		if 0 < fullGame.MoveCount {
			return nil, errors.New("Black player has already played, and cannot reject")
		}
	} else {
		return nil, errors.New("Message creator is not a player")
	}

	// Remove the game completely as it is not interesting to keep it.
	k.Keeper.RemoveStoredGame(ctx, msg.IdValue)

	// What to emit
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.RejectGameEventKey,
			sdk.NewAttribute(types.RejectGameEventCreator, msg.Creator),
			sdk.NewAttribute(types.RejectGameEventIdValue, msg.IdValue),
		),
	)

	return &types.MsgRejectGameResponse{}, nil
}
