package keeper

import (
	"context"
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
	newGame := types.FullGame{
		Creator: sdk.AccAddress(msg.Creator),
		Index:   newIndex,
		Game:    *rules.New(),
		Red:     sdk.AccAddress(msg.Red),
		Black:   sdk.AccAddress(msg.Black),
	}
	k.Keeper.SetStoredGame(ctx, *newGame.ToStoredGame())

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	// What to emit
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.StoredGameEventKey,
			sdk.NewAttribute(types.StoredGameEventCreator, msg.Creator),
			sdk.NewAttribute(types.StoredGameEventIndex, newIndex),
			sdk.NewAttribute(types.StoredGameEventRed, msg.Red),
			sdk.NewAttribute(types.StoredGameEventBlack, msg.Black),
		),
	)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
