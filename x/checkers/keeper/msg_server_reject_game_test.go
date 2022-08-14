package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/testutil/mock_types"
	"github.com/b9lab/checkers/x/checkers"
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMsgServerWithOneGameForRejectGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *mock_types.MockBankEscrowKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := mock_types.NewMockBankEscrowKeeper(ctrl)
	k, ctx := keepertest.CheckersKeeperWithMocks(t, bankMock)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
	return server, *k, context, ctrl, bankMock
}

func TestRejectGameWrongByCreator(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	defer ctrl.Finish()
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   alice,
		GameIndex: "1",
	})
	require.Nil(t, rejectGameResponse)
	require.Equal(t, alice+": message creator is not a player", err.Error())
}

func TestRejectGameByBlackNoMove(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	defer ctrl.Finish()
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   bob,
		GameIndex: "1",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func TestRejectGameByBlackNoMoveRemovedGame(t *testing.T) {
	msgServer, keeper, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   bob,
		GameIndex: "1",
	})
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)
	_, found = keeper.GetStoredGame(ctx, "1")
	require.False(t, found)
}

func TestRejectGameByBlackNoMoveEmitted(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   bob,
		GameIndex: "1",
	})
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 2)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "game-rejected",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: bob},
			{Key: "game-index", Value: "1"},
		},
	}, event)
}

func TestRejectGameByBlackRefundedGas(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	before := ctx.GasMeter().GasConsumed()
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   bob,
		GameIndex: "1",
	})
	after := ctx.GasMeter().GasConsumed()
	require.LessOrEqual(t, after, before-5_000)
}

func TestRejectGameByRedNoMove(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	defer ctrl.Finish()
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func TestRejectGameByRedNoMoveRemovedGame(t *testing.T) {
	msgServer, keeper, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)
	_, found = keeper.GetStoredGame(ctx, "1")
	require.False(t, found)
}

func TestRejectGameByRedNoMoveEmitted(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 2)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "game-rejected",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: carol},
			{Key: "game-index", Value: "1"},
		},
	}, event)
}

func TestRejectGameByRedOneMove(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForRejectGame(t)
	defer ctrl.Finish()
	escrow.ExpectAny(context)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func TestRejectGameByRedOneMoveRemovedGame(t *testing.T) {
	msgServer, keeper, context, ctrl, escrow := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectAny(context)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "-1",
		FifoTailIndex: "-1",
	}, systemInfo)
	_, found = keeper.GetStoredGame(ctx, "1")
	require.False(t, found)
}

func TestRejectGameByRedOneMoveEmitted(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForRejectGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectAny(context)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 3)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "game-rejected",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: carol},
			{Key: "game-index", Value: "1"},
		},
	}, event)
}

func TestRejectGameByRedOneCalledBank(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForRejectGame(t)
	defer ctrl.Finish()
	payBob := escrow.ExpectPay(context, bob, 45).Times(1)
	escrow.ExpectRefund(context, bob, 45).Times(1).After(payBob)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
}

func TestRejectGameByBlackWrongOneMove(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForRejectGame(t)
	defer ctrl.Finish()
	escrow.ExpectPay(context, bob, 45).Times(1)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   bob,
		GameIndex: "1",
	})
	require.Nil(t, rejectGameResponse)
	require.Equal(t, "black player has already played", err.Error())
}

func TestRejectGameByRedWrong2Moves(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForRejectGame(t)
	defer ctrl.Finish()
	payBob := escrow.ExpectPay(context, bob, 45).Times(1)
	escrow.ExpectPay(context, carol, 45).Times(1).After(payBob)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	rejectGameResponse, err := msgServer.RejectGame(context, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	require.Nil(t, rejectGameResponse)
	require.Equal(t, "red player has already played", err.Error())
}
