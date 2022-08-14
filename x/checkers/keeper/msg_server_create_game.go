package keeper

import (
	"context"
	"strconv"

	"github.com/b9lab/checkers/x/checkers/rules"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	newGame := rules.New()
	storedGame := types.StoredGame{
		Index:       newIndex,
		Board:       newGame.String(),
		Turn:        rules.PieceStrings[newGame.Turn],
		Black:       msg.Black,
		Red:         msg.Red,
		MoveCount:   0,
		BeforeIndex: types.NoFifoIndex,
		AfterIndex:  types.NoFifoIndex,
		Deadline:    types.FormatDeadline(types.GetNextDeadline(ctx)),
		Winner:      rules.PieceStrings[rules.NO_PLAYER],
		Wager:       msg.Wager,
	}

	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SendToFifoTail(ctx, &storedGame, &systemInfo)
	k.Keeper.SetStoredGame(ctx, storedGame)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	ctx.GasMeter().ConsumeGas(types.CreateGameGas, "Create game")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, newIndex),
			sdk.NewAttribute(types.GameCreatedEventBlack, msg.Black),
			sdk.NewAttribute(types.GameCreatedEventRed, msg.Red),
			sdk.NewAttribute(types.GameCreatedEventWager, strconv.FormatUint(msg.Wager, 10)),
		),
	)

	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}
