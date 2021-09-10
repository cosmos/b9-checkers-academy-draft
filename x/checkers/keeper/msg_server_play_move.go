package keeper

import (
	"context"
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rules "github.com/xavierlepretre/checkers/x/checkers/rules"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.IdValue)
	if !found {
		return nil, errors.New("Game not found " + msg.IdValue)
	}

	// Is it an expected player?
	var player rules.Player
	if strings.Compare(storedGame.Red, msg.Creator) == 0 {
		player = rules.RED_PLAYER
	} else if strings.Compare(storedGame.Black, msg.Creator) == 0 {
		player = rules.BLACK_PLAYER
	} else {
		return nil, errors.New("Message creator is not a player")
	}

	// Is it the player's turn?
	fullGame := storedGame.ToFullGame()
	if !fullGame.Game.TurnIs(player) {
		return nil, errors.New("Player tried to play out of turn")
	}

	// Do it
	captured, moveErr := fullGame.Game.Move(
		rules.Pos{
			X: int(msg.FromX),
			Y: int(msg.FromY),
		},
		rules.Pos{
			X: int(msg.ToX),
			Y: int(msg.ToY),
		},
	)
	if moveErr != nil {
		return nil, moveErr
	}

	// Save for the next play move
	k.Keeper.SetStoredGame(ctx, *fullGame.ToStoredGame())

	// What to inform
	return &types.MsgPlayMoveResponse{
		IdValue:   msg.IdValue,
		CapturedX: int64(captured.X),
		CapturedY: int64(captured.Y),
		Winner:    fullGame.Game.Winner().Color,
	}, nil
}
