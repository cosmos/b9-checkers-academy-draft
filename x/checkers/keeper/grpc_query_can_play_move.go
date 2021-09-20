package keeper

import (
	"context"
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rules "github.com/xavierlepretre/checkers/x/checkers/rules"
	"github.com/xavierlepretre/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CanPlayMove(goCtx context.Context, req *types.QueryCanPlayMoveRequest) (*types.QueryCanPlayMoveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.GetStoredGame(ctx, req.IdValue)
	if !found {
		return nil, errors.New("Game not found " + req.IdValue)
	}

	// Is the game already won? Can happen when it is forfeited.
	if storedGame.Winner != rules.NO_PLAYER.Color {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   "Game is already finished",
		}, nil
	}

	// Is it an expected player?
	var player rules.Player
	if strings.Compare(rules.RED_PLAYER.Color, req.Player) == 0 {
		player = rules.RED_PLAYER
	} else if strings.Compare(rules.BLACK_PLAYER.Color, req.Player) == 0 {
		player = rules.BLACK_PLAYER
	} else {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   "Message creator is not a player",
		}, nil
	}

	// Is it the player's turn?
	fullGame := storedGame.ToFullGame()
	if !fullGame.Game.TurnIs(player) {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   "Player tried to play out of turn",
		}, nil
	}

	// Attempt a move in memory
	_, moveErr := fullGame.Game.Move(
		rules.Pos{
			X: int(req.FromX),
			Y: int(req.FromY),
		},
		rules.Pos{
			X: int(req.ToX),
			Y: int(req.ToY),
		},
	)
	if moveErr != nil {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   moveErr.Error(),
		}, nil
	}

	return &types.QueryCanPlayMoveResponse{
		Possible: true,
		Reason:   "Ok",
	}, nil
}
