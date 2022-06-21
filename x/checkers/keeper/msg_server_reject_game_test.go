package keeper_test

import (
	"context"
	"testing"

	"github.com/b9lab/checkers/x/checkers"
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerWithOneGameForRejectGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
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

func TestRejectGameWrongByCreator(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: alice,
		IdValue: "1",
	})
	require.Nil(t, rejectGameResponse)
	require.Equal(t, "message creator is not a player: %s", err.Error())
}

func TestRejectGameByBlackNoMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func TestRejectGameByBlackNoMoveRemovedGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	nextGame, found := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	_, found = keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.False(t, found)
}

func TestRejectGameByBlackNoMoveEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.Equal(t, event.Type, "message")
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameRejected"},
		{Key: "Creator", Value: carol},
		{Key: "IdValue", Value: "1"},
	}, event.Attributes[7:])
}

func TestRejectGameByRedNoMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func TestRejectGameByRedNoMoveRemovedGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	nextGame, found := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	_, found = keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.False(t, found)
}

func TestRejectGameByRedNoMoveEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.Equal(t, event.Type, "message")
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameRejected"},
		{Key: "Creator", Value: bob},
		{Key: "IdValue", Value: "1"},
	}, event.Attributes[7:])
}

func TestRejectGameByRedOneMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func TestRejectGameByRedOneMoveRemovedGame(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	nextGame, found := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	_, found = keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.False(t, found)
}

func TestRejectGameByRedOneMoveEmitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.Equal(t, event.Type, "message")
	require.EqualValues(t, []sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameRejected"},
		{Key: "Creator", Value: bob},
		{Key: "IdValue", Value: "1"},
	}, event.Attributes[14:])
}

func TestRejectGameByBlackWrongOneMove(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	require.Nil(t, rejectGameResponse)
	require.Equal(t, "black player has already played", err.Error())
}

func TestRejectGameByRedWrong2Moves(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameForRejectGame(t)
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
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	require.Nil(t, rejectGameResponse)
	require.Equal(t, "red player has already played", err.Error())
}
