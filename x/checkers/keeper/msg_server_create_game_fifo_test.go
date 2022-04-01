package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func TestCreate3GamesHasSavedFifo(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
	})

	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
	})
	nextGame2, found2 := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found2)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  3,
		FifoHead: "1",
		FifoTail: "2",
	}, nextGame2)
	game1, found1 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "2",
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
		BeforeId:  "1",
		AfterId:   "-1",
	}, game2)

	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
	})
	nextGame3, found3 := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found3)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  4,
		FifoHead: "1",
		FifoTail: "3",
	}, nextGame3)
	game1, found1 = keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "2",
	}, game1)
	game2, found2 = keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "2")
	require.True(t, found2)
	require.EqualValues(t, types.StoredGame{
		Creator:   bob,
		Index:     "2",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       carol,
		Black:     alice,
		MoveCount: uint64(0),
		BeforeId:  "1",
		AfterId:   "3",
	}, game2)
	game3, found3 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "3")
	require.True(t, found3)
	require.EqualValues(t, types.StoredGame{
		Creator:   carol,
		Index:     "3",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       alice,
		Black:     bob,
		MoveCount: uint64(0),
		BeforeId:  "2",
		AfterId:   "-1",
	}, game3)
}
