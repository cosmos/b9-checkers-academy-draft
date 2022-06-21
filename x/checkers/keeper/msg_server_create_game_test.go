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

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := setupKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		IdValue: "1",
	}, *createResponse)
}

func TestCreate1GameHasSaved(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
	})
	nextGame, found := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextGame)
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
		AfterId:   "-1",
	}, game1)
}

func TestCreate1GameGetAll(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
	})
	games := keeper.GetAllStoredGame(sdk.UnwrapSDKContext(context))
	require.Len(t, games, 1)
	require.EqualValues(t, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "-1",
	}, games[0])
}

func TestCreate1GameEmitted(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "message",
		Attributes: []sdk.Attribute{
			{Key: "module", Value: "checkers"},
			{Key: "action", Value: "NewGameCreated"},
			{Key: "Creator", Value: alice},
			{Key: "Index", Value: "1"},
			{Key: "Red", Value: bob},
			{Key: "Black", Value: carol},
		},
	}, event)
}

func TestCreateGameRedAddressBad(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     "notanaddress",
		Black:   carol,
	})
	require.Nil(t, createResponse)
	require.Equal(t,
		"red address is invalid: notanaddress: decoding bech32 failed: invalid index of 1",
		err.Error())
}

func TestCreateGameEmptyRedAddress(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     "",
		Black:   carol,
	})
	require.Nil(t, createResponse)
	require.Equal(t,
		"red address is invalid: : empty address string is not allowed",
		err.Error())
}

func TestCreate3Games(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
	})
	createResponse2, err2 := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
	})
	require.Nil(t, err2)
	require.EqualValues(t, types.MsgCreateGameResponse{
		IdValue: "2",
	}, *createResponse2)
	createResponse3, err3 := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
	})
	require.Nil(t, err3)
	require.EqualValues(t, types.MsgCreateGameResponse{
		IdValue: "3",
	}, *createResponse3)
}

func TestCreate3GamesHasSaved(t *testing.T) {
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
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
	})
	nextGame, found := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		Creator:  "",
		IdValue:  4,
		FifoHead: "1",
		FifoTail: "3",
	}, nextGame)
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

func TestCreate3GamesGetAll(t *testing.T) {
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
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
	})
	games := keeper.GetAllStoredGame(sdk.UnwrapSDKContext(context))
	require.Len(t, games, 3)
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
	}, games[0])
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
	}, games[1])
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
	}, games[2])
}
