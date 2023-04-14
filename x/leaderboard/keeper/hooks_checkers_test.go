package keeper_test

import (
	"sort"
	"testing"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	checkerstypes "github.com/b9lab/checkers/x/checkers/types"
	"github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestOneCandidateAdded(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	keeper.Hooks().AfterPlayerInfoChanged(ctx, checkerstypes.PlayerInfo{
		Index:          alice,
		WonCount:       12,
		LostCount:      13,
		ForfeitedCount: 14,
	})

	aliceAddress, err := sdk.AccAddressFromBech32(alice)
	require.Nil(t, err)
	candidates := keeper.GetAllCandidates(ctx)
	require.Len(t, candidates, 1)
	require.Equal(t,
		types.Candidate{Address: aliceAddress, WonCount: 12},
		candidates[0],
	)
}

func TestOneCandidateOverwritten(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	keeper.Hooks().AfterPlayerInfoChanged(ctx, checkerstypes.PlayerInfo{
		Index:          alice,
		WonCount:       12,
		LostCount:      13,
		ForfeitedCount: 14,
	})
	keeper.Hooks().AfterPlayerInfoChanged(ctx, checkerstypes.PlayerInfo{
		Index:          alice,
		WonCount:       22,
		LostCount:      23,
		ForfeitedCount: 24,
	})

	aliceAddress, err := sdk.AccAddressFromBech32(alice)
	require.Nil(t, err)
	candidates := keeper.GetAllCandidates(ctx)
	require.Len(t, candidates, 1)
	require.Equal(t,
		types.Candidate{Address: aliceAddress, WonCount: 22},
		candidates[0],
	)
}

func TestOneCandidateAddedAlongside(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	aliceAddress, err := sdk.AccAddressFromBech32(alice)
	require.Nil(t, err)
	keeper.SetCandidate(ctx, types.Candidate{Address: aliceAddress, WonCount: 12})
	keeper.Hooks().AfterPlayerInfoChanged(ctx, checkerstypes.PlayerInfo{
		Index:          bob,
		WonCount:       22,
		LostCount:      23,
		ForfeitedCount: 24,
	})

	candidates := keeper.GetAllCandidates(ctx)
	require.Len(t, candidates, 2)
	sort.SliceStable(candidates[:], func(i, j int) bool {
		return candidates[i].WonCount < candidates[j].WonCount
	})
	bobAddress, err := sdk.AccAddressFromBech32(bob)
	require.Nil(t, err)
	require.Equal(t,
		[]types.Candidate{
			{Address: aliceAddress, WonCount: 12},
			{Address: bobAddress, WonCount: 22},
		},
		candidates,
	)
}
