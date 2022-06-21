package keeper_test

import (
	"testing"

	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestPlayMove2Games1MoveHasSavedFifo(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
	})

	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	nextGame1, found1 := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found1)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  3,
		FifoHead: "2",
		FifoTail: "1",
	}, nextGame1)
	game1, found1 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "r",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(1),
		BeforeId:  "2",
		AfterId:   "-1",
	}, game1)
	game2, found2 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "2")
	require.True(t, found2)
	require.EqualValues(t, types.StoredGame{
		Creator:   bob,
		Index:     "2",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       carol,
		Black:     alice,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "1",
	}, game2)
}

func TestPlayMove2Games2MovesHasSavedFifo(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
	})
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})

	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator: alice,
		IdValue: "2",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	nextGame1, found1 := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found1)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  3,
		FifoHead: "1",
		FifoTail: "2",
	}, nextGame1)
	game1, found1 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "r",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(1),
		BeforeId:  "-1",
		AfterId:   "2",
	}, game1)
	game2, found2 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "2")
	require.True(t, found2)
	require.EqualValues(t, types.StoredGame{
		Creator:   bob,
		Index:     "2",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "r",
		Red:       carol,
		Black:     alice,
		MoveCount: uint64(1),
		BeforeId:  "1",
		AfterId:   "-1",
	}, game2)
}
