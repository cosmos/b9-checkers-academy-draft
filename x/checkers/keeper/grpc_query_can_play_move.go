package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, types.ErrGameNotFound.Error(), req.IdValue)
	}

	// Is the game already won? Can happen when it is forfeited.
	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   types.ErrGameFinished.Error(),
		}, nil
	}

	// Is it an expected player?
	var player rules.Player
	if strings.Compare(rules.PieceStrings[rules.RED_PLAYER], req.Player) == 0 {
		player = rules.RED_PLAYER
	} else if strings.Compare(rules.PieceStrings[rules.BLACK_PLAYER], req.Player) == 0 {
		player = rules.BLACK_PLAYER
	} else {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   fmt.Sprintf(types.ErrCreatorNotPlayer.Error(), req.Player),
		}, nil
	}

	// Is it the player's turn?
	game, err := storedGame.ParseGame()
	if err != nil {
		return nil, err
	}
	if !game.TurnIs(player) {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   fmt.Sprintf(types.ErrNotPlayerTurn.Error(), player.Color),
		}, nil
	}

	// Attempt a move in memory
	_, moveErr := game.Move(
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
			Reason:   fmt.Sprintf(types.ErrWrongMove.Error(), moveErr.Error()),
		}, nil
	}

	return &types.QueryCanPlayMoveResponse{
		Possible: true,
		Reason:   "ok",
	}, nil
}
