package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/xavierlepretre/checkers/x/checkers/keeper"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := setupKeeper(t)
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
		IdValue: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
