package keeper_test

import (
	"context"
	"errors"
	"testing"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/x/checkers"
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/testutil"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupKeeperForWagerHandler(t testing.TB) (keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankEscrowKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankEscrowKeeper(ctrl)
	k, ctx := keepertest.CheckersKeeperWithMocks(t, bankMock)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	context := sdk.WrapSDKContext(ctx)
	return *k, context, ctrl, bankMock
}

func TestWagerHandlerCollectWrongNoBlack(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "black address is invalid: : empty address string is not allowed", r)
	}()
	keeper.CollectWager(ctx, &types.StoredGame{
		MoveCount: 0,
	})
}

func TestWagerHandlerCollectFailedNoMove(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	black, _ := sdk.AccAddressFromBech32(alice)
	escrow.EXPECT().
		SendCoinsFromAccountToModule(ctx, black, types.ModuleName, gomock.Any()).
		Return(errors.New("Oops"))
	err := keeper.CollectWager(ctx, &types.StoredGame{
		Black:     alice,
		MoveCount: 0,
		Wager:     45,
	})
	require.NotNil(t, err)
	require.EqualError(t, err, "black cannot pay the wager: Oops")
}

func TestWagerHandlerCollectWrongNoRed(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "red address is invalid: : empty address string is not allowed", r)
	}()
	keeper.CollectWager(ctx, &types.StoredGame{
		MoveCount: 1,
	})
}

func TestWagerHandlerCollectFailedOneMove(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	red, _ := sdk.AccAddressFromBech32(bob)
	escrow.EXPECT().
		SendCoinsFromAccountToModule(ctx, red, types.ModuleName, gomock.Any()).
		Return(errors.New("Oops"))
	err := keeper.CollectWager(ctx, &types.StoredGame{
		Red:       bob,
		MoveCount: 1,
		Wager:     45,
	})
	require.NotNil(t, err)
	require.EqualError(t, err, "red cannot pay the wager: Oops")
}

func TestWagerHandlerCollectNoMove(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectPay(context, alice, 45)
	err := keeper.CollectWager(ctx, &types.StoredGame{
		Black:     alice,
		MoveCount: 0,
		Wager:     45,
	})
	require.Nil(t, err)
}

func TestWagerHandlerCollectOneMove(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectPay(context, bob, 45)
	err := keeper.CollectWager(ctx, &types.StoredGame{
		Red:       bob,
		MoveCount: 1,
		Wager:     45,
	})
	require.Nil(t, err)
}

func TestWagerHandlerPayWrongNoWinnerAddress(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "black address is invalid: : empty address string is not allowed", r)
	}()
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Winner: "b",
	})
}

func TestWagerHandlerPayWrongWinnerNotFound(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "cannot find winner by color: *", r)
	}()
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Black:  alice,
		Red:    bob,
		Winner: "*",
	})
}

func TestWagerHandlerPayWrongNotPayTime(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "there is nothing to pay, should not have been called", r)
	}()
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Black:     alice,
		Red:       bob,
		Winner:    "b",
		MoveCount: 0,
	})
}

func TestWagerHandlerPayWrongEscrowFailed(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	black, _ := sdk.AccAddressFromBech32(alice)
	escrow.EXPECT().
		SendCoinsFromModuleToAccount(ctx, types.ModuleName, black, gomock.Any()).
		Times(1).
		Return(errors.New("Oops"))
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, r, "cannot pay winnings to winner: Oops")
	}()
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Black:     alice,
		Red:       bob,
		Winner:    "b",
		MoveCount: 1,
		Wager:     45,
	})
}

func TestWagerHandlerPayEscrowCalledOneMove(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectRefund(context, alice, 45)
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Black:     alice,
		Red:       bob,
		Winner:    "b",
		MoveCount: 1,
		Wager:     45,
	})
}

func TestWagerHandlerPayEscrowCalledTwoMoves(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectRefund(context, alice, 90)
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Black:     alice,
		Red:       bob,
		Winner:    "b",
		MoveCount: 2,
		Wager:     45,
	})
}

func TestWagerHandlerRefundWrongManyMoves(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "game is not in a state to refund, move count: 2", r)
	}()
	keeper.MustRefundWager(ctx, &types.StoredGame{
		MoveCount: 2,
	})
}

func TestWagerHandlerRefundNoMoves(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	keeper.MustRefundWager(ctx, &types.StoredGame{
		MoveCount: 0,
	})
}

func TestWagerHandlerRefundWrongNoBlack(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "black address is invalid: : empty address string is not allowed", r)
	}()
	keeper.MustRefundWager(ctx, &types.StoredGame{
		MoveCount: 1,
	})
}

func TestWagerHandlerRefundWrongEscrowFailed(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	black, _ := sdk.AccAddressFromBech32(alice)
	escrow.EXPECT().
		SendCoinsFromModuleToAccount(ctx, types.ModuleName, black, gomock.Any()).
		Times(1).
		Return(errors.New("Oops"))
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "cannot refund wager to: Oops", r)
	}()
	keeper.MustRefundWager(ctx, &types.StoredGame{
		Black:     alice,
		MoveCount: 1,
		Wager:     45,
	})
}

func TestWagerHandlerRefundCalled(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectRefund(context, alice, 45)
	keeper.MustRefundWager(ctx, &types.StoredGame{
		Black:     alice,
		MoveCount: 1,
		Wager:     45,
	})
}
