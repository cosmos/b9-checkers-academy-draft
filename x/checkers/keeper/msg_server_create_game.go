package keeper

import (
	"context"
	"errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rules "github.com/xavierlepretre/checkers/x/checkers/rules"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	newIndex := strconv.FormatUint(nextGame.IdValue, 10)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.New("Creator is a malformed address")
	}
	red, err := sdk.AccAddressFromBech32(msg.Red)
	if err != nil {
		return nil, errors.New("Red is a malformed address")
	}
	black, err := sdk.AccAddressFromBech32(msg.Black)
	if err != nil {
		return nil, errors.New("Black is a malformed address")
	}
	newGame := types.FullGame{
		Creator: creator,
		Index:   newIndex,
		Game:    *rules.New(),
		Red:     red,
		Black:   black,
	}
	k.Keeper.SetStoredGame(ctx, newGame.ToStoredGame())

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
