package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
