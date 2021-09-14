package keeper

import (
	"context"
	"errors"
	"strconv"
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

	// Is the game already won? Can happen when it is forfeited.
	if storedGame.Winner != rules.NO_PLAYER.Color {
		return nil, errors.New("Game is already finished")
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
	fullGame.MoveCount++
	fullGame.Deadline = ctx.BlockTime().Add(types.MaxTurnDurationInSeconds)
	fullGame.Winner = fullGame.Game.Winner().Color

	// Remove from or send to the back of the FIFO
	storedGame = *fullGame.ToStoredGame()
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	if storedGame.Winner == rules.NO_PLAYER.Color {
		k.Keeper.SendToFifoTail(ctx, &storedGame, &nextGame)
	} else {
		k.Keeper.RemoveFromFifo(ctx, &storedGame, &nextGame)
	}

	// Save for the next play move
	k.Keeper.SetStoredGame(ctx, storedGame)
	k.Keeper.SetNextGame(ctx, nextGame)

	// What to emit
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PlayMoveEventKey,
			sdk.NewAttribute(types.PlayMoveEventCreator, msg.Creator),
			sdk.NewAttribute(types.PlayMoveEventIdValue, msg.IdValue),
			sdk.NewAttribute(types.PlayMoveEventCapturedX, strconv.FormatInt(int64(captured.X), 10)),
			sdk.NewAttribute(types.PlayMoveEventCapturedY, strconv.FormatInt(int64(captured.Y), 10)),
			sdk.NewAttribute(types.PlayMoveEventWinner, fullGame.Winner),
		),
	)

	// What to inform
	return &types.MsgPlayMoveResponse{
		IdValue:   msg.IdValue,
		CapturedX: int64(captured.X),
		CapturedY: int64(captured.Y),
		Winner:    fullGame.Winner,
	}, nil
}
