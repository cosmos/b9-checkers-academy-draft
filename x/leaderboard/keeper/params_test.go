package keeper_test

import (
	"testing"

	testkeeper "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/x/leaderboard/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LeaderboardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Length, k.Length(ctx))
}
