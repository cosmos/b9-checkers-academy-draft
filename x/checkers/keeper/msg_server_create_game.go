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
		Creator:  sdk.AccAddress(msg.Creator),
		Index:    newIndex,
		Game:     *rules.New(),
		Red:      sdk.AccAddress(msg.Red),
		Black:    sdk.AccAddress(msg.Black),
		Deadline: ctx.BlockTime().Add(types.MaxTurnDurationInSeconds),
		Winner:   rules.NO_PLAYER.Color,
		Wager:    sdk.NewCoin(msg.Token, sdk.NewInt(int64(msg.Wager))),
	}
	storedGame := newGame.ToStoredGame()
	k.Keeper.SendToFifoTail(ctx, &storedGame, &nextGame)
	k.Keeper.SetStoredGame(ctx, storedGame)

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	ctx.GasMeter().ConsumeGas(types.CreateGameGas, "Create game")

	// What to emit
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.StoredGameEventKey,
			sdk.NewAttribute(types.StoredGameEventCreator, msg.Creator),
			sdk.NewAttribute(types.StoredGameEventIndex, newIndex),
			sdk.NewAttribute(types.StoredGameEventRed, msg.Red),
			sdk.NewAttribute(types.StoredGameEventBlack, msg.Black),
			sdk.NewAttribute(types.StoredGameEventWager, strconv.FormatUint(msg.Wager, 10)),
			sdk.NewAttribute(types.StoredGameEventToken, msg.Token),
		),
	)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
