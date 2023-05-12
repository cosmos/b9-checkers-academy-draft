package keeper_test

import (
	"testing"

	"github.com/b9lab/checkers/x/checkers/testutil"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestPlayMoveUpToWinner(t *testing.T) {
	msgServer, keeper, context, ctrl, escrow := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectAny(context)

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)

	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: types.NoFifoIndex,
		FifoTailIndex: types.NoFifoIndex,
	}, systemInfo)

	game, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "b",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   40,
		BeforeIndex: types.NoFifoIndex,
		AfterIndex:  types.NoFifoIndex,
		Wager:       45,
		Denom:       "stake",
	}, game)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 2)
	event := events[0]
	require.Equal(t, event.Type, "move-played")
	require.EqualValues(t, []sdk.Attribute{
		{Key: "creator", Value: bob},
		{Key: "game-index", Value: "1"},
		{Key: "captured-x", Value: "2"},
		{Key: "captured-y", Value: "5"},
		{Key: "winner", Value: "b"},
		{Key: "board", Value: "*b*b****|**b*b***|*****b**|********|***B****|********|*****b**|********"},
	}, event.Attributes[(len(testutil.Game1Moves)-1)*6:])
}

func TestPlayMoveUpToWinnerCalledBank(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForPlayMove(t)
	defer ctrl.Finish()
	payBob := escrow.ExpectPay(context, bob, 45).Times(1)
	payCarol := escrow.ExpectPay(context, carol, 45).Times(1).After(payBob)
	escrow.ExpectRefund(context, bob, 90).Times(1).After(payCarol)

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)
}

func TestCompleteGameAddPlayerInfo(t *testing.T) {
	msgServer, keeper, context, ctrl, escrow := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectAny(context)

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)

	bobInfo, found := keeper.GetPlayerInfo(ctx, bob)
	require.True(t, found)
	require.EqualValues(t, types.PlayerInfo{
		Index:          bob,
		WonCount:       1,
		LostCount:      0,
		ForfeitedCount: 0,
	}, bobInfo)
	carolInfo, found := keeper.GetPlayerInfo(ctx, carol)
	require.True(t, found)
	require.EqualValues(t, types.PlayerInfo{
		Index:          carol,
		WonCount:       0,
		LostCount:      1,
		ForfeitedCount: 0,
	}, carolInfo)
}

func TestCompleteGameUpdatePlayerInfo(t *testing.T) {
	msgServer, keeper, context, ctrl, escrow := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectAny(context)

	keeper.SetPlayerInfo(ctx, types.PlayerInfo{
		Index:          bob,
		WonCount:       1,
		LostCount:      2,
		ForfeitedCount: 3,
	})
	keeper.SetPlayerInfo(ctx, types.PlayerInfo{
		Index:          carol,
		WonCount:       4,
		LostCount:      5,
		ForfeitedCount: 6,
	})

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)

	bobInfo, found := keeper.GetPlayerInfo(ctx, bob)
	require.True(t, found)
	require.EqualValues(t, types.PlayerInfo{
		Index:          bob,
		WonCount:       2,
		LostCount:      2,
		ForfeitedCount: 3,
	}, bobInfo)
	carolInfo, found := keeper.GetPlayerInfo(ctx, carol)
	require.True(t, found)
	require.EqualValues(t, types.PlayerInfo{
		Index:          carol,
		WonCount:       4,
		LostCount:      6,
		ForfeitedCount: 6,
	}, carolInfo)
}

func TestCompleteGameCallsHook(t *testing.T) {
	msgServer, keeper, context, ctrl, escrow, hookMock := setupMsgServerWithOneGameForPlayMoveAndHooks(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectAny(context)
	bobCall := hookMock.EXPECT().AfterPlayerInfoChanged(ctx, types.PlayerInfo{
		Index:          bob,
		WonCount:       2,
		LostCount:      2,
		ForfeitedCount: 3,
	}).Times(1)
	hookMock.EXPECT().AfterPlayerInfoChanged(ctx, types.PlayerInfo{
		Index:          carol,
		WonCount:       4,
		LostCount:      6,
		ForfeitedCount: 6,
	}).Times(1).After(bobCall)

	keeper.SetPlayerInfo(ctx, types.PlayerInfo{
		Index:          bob,
		WonCount:       1,
		LostCount:      2,
		ForfeitedCount: 3,
	})
	keeper.SetPlayerInfo(ctx, types.PlayerInfo{
		Index:          carol,
		WonCount:       4,
		LostCount:      5,
		ForfeitedCount: 6,
	})

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)
}
