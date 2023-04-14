package keeper_test

import (
	"testing"
	"time"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestOnePlayerAddedToLeaderboard(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	keeper.SetLeaderboard(ctx, types.Leaderboard{
		Winners: []types.Winner{
			{Address: alice, WonCount: 12, AddedAt: 999},
			{Address: carol, WonCount: 10, AddedAt: 999},
		},
	})
	bobAddress, err := sdk.AccAddressFromBech32(bob)
	require.Nil(t, err)
	keeper.SetCandidate(ctx, types.Candidate{
		Address:  bobAddress,
		WonCount: 11,
	})

	bobTime := time.Unix(1000, 0)
	keeper.CollectSortAndClipLeaderboard(ctx.WithBlockTime(bobTime))

	leaderboard := keeper.GetLeaderboard(ctx)
	require.Len(t, leaderboard.Winners, 3)
	require.Equal(t,
		[]types.Winner{
			{Address: alice, WonCount: 12, AddedAt: 999},
			{Address: bob, WonCount: 11, AddedAt: 1000},
			{Address: carol, WonCount: 10, AddedAt: 999},
		},
		leaderboard.Winners,
	)
}

func TestOnePlayerAddedAndOneUpdatedToLeaderboard(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	keeper.SetLeaderboard(ctx, types.Leaderboard{
		Winners: []types.Winner{
			{Address: alice, WonCount: 12, AddedAt: 999},
			{Address: bob, WonCount: 10, AddedAt: 999},
		},
	})
	bobAddress, err := sdk.AccAddressFromBech32(bob)
	require.Nil(t, err)
	carolAddress, err := sdk.AccAddressFromBech32(carol)
	require.Nil(t, err)
	keeper.SetCandidate(ctx, types.Candidate{
		Address:  bobAddress,
		WonCount: 13,
	})
	keeper.SetCandidate(ctx, types.Candidate{
		Address:  carolAddress,
		WonCount: 12,
	})
	bobTime := time.Unix(1000, 0)
	keeper.CollectSortAndClipLeaderboard(ctx.WithBlockTime(bobTime))

	leaderboard := keeper.GetLeaderboard(ctx)
	require.Len(t, leaderboard.Winners, 3)
	require.Equal(t,
		[]types.Winner{
			{Address: bob, WonCount: 13, AddedAt: 1000},
			{Address: carol, WonCount: 12, AddedAt: 1000},
			{Address: alice, WonCount: 12, AddedAt: 999},
		},
		leaderboard.Winners,
	)
}

func TestOnePlayerKicksPlayerOutOfLeaderboard(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	keeper.SetLeaderboard(ctx, types.Leaderboard{
		Winners: []types.Winner{
			{Address: alice, WonCount: 12, AddedAt: 999},
			{Address: bob, WonCount: 10, AddedAt: 999},
		},
	})
	params := keeper.GetParams(ctx)
	params.Length = 2
	keeper.SetParams(ctx, params)
	carolAddress, err := sdk.AccAddressFromBech32(carol)
	require.Nil(t, err)
	keeper.SetCandidate(ctx, types.Candidate{
		Address:  carolAddress,
		WonCount: 11,
	})
	carolTime := time.Unix(1000, 0)
	keeper.CollectSortAndClipLeaderboard(ctx.WithBlockTime(carolTime))

	leaderboard := keeper.GetLeaderboard(ctx)
	require.Len(t, leaderboard.Winners, 2)
	require.Equal(t,
		[]types.Winner{
			{Address: alice, WonCount: 12, AddedAt: 999},
			{Address: carol, WonCount: 11, AddedAt: 1000},
		},
		leaderboard.Winners,
	)
}
