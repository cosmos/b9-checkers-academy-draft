package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/b9blab/checkers/testutil/keeper"
	"github.com/b9blab/checkers/x/checkers/keeper"
	"github.com/b9blab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	require.NotNil(t, msgServer)
	require.NotNil(t, context)
}
