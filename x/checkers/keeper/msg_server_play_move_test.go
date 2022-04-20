package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/xavierlepretre/checkers/x/checkers"
	"github.com/xavierlepretre/checkers/x/checkers/keeper"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func setupMsgServerWithOneGameForPlayMove(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := setupKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   11,
	})
	return server, *k, context
}

func TestPlayMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
	playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func TestPlayMoveSameBlackRed(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   bob,
	})
	playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func TestPlayMoveSavedGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	nextGame, found := keeper.GetNextGame(ctx)
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextGame)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "r",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(1),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     11,
	}, game1)
}

func TestPlayMoveWrongOutOfTurn(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
	playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	require.Nil(t, playMoveResponse)
	require.Equal(t, "player tried to play out of turn: %s", err.Error())
}

func TestPlayMoveWrongPieceAtDestination(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
	playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   0,
		ToX:     0,
		ToY:     1,
	})
	require.Nil(t, playMoveResponse)
	require.Equal(t, "Already piece at destination position: {0 1}: wrong move", err.Error())
}

func TestPlayMoveEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.Equal(t, event.Type, "message")
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "MovePlayed"},
		{Key: "Creator", Value: carol},
		{Key: "IdValue", Value: "1"},
		{Key: "CapturedX", Value: "-1"},
		{Key: "CapturedY", Value: "-1"},
		{Key: "Winner", Value: "*"},
	}, event.Attributes[7:])
}

func TestPlayMove2(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func TestPlayMove2SavedGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	nextGame, found := keeper.GetNextGame(ctx)
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextGame)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(2),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     11,
	}, game1)
}

func TestPlayMove3(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   2,
		FromY:   3,
		ToX:     0,
		ToY:     5,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: 1,
		CapturedY: 4,
		Winner:    "*",
	}, *playMoveResponse)
}

func TestPlayMove3SavedGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   2,
		FromY:   3,
		ToX:     0,
		ToY:     5,
	})
	nextGame, found := keeper.GetNextGame(ctx)
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextGame)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|********|********|b*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "r",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(3),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     11,
	}, game1)
}
