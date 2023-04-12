package keeper_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/x/checkers"
	"github.com/b9lab/checkers/x/checkers/keeper"
	cv3keeper "github.com/b9lab/checkers/x/checkers/migrations/cv3/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupKeeperForV1ToV1_1Migration(t testing.TB) (keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return *k, sdk.WrapSDKContext(ctx)
}

func TestBuildPlayerInfosInPlace(t *testing.T) {
	tests := []struct {
		name     string
		games    []types.StoredGame
		expected []types.PlayerInfo
	}{
		{
			name:     "nothing to assemble",
			games:    []types.StoredGame{},
			expected: []types.PlayerInfo(nil),
		},
		{
			name: "single game no win",
			games: []types.StoredGame{
				{
					Index:  "1",
					Winner: "*",
					Black:  "alice",
					Red:    "bob",
				},
			},
			expected: []types.PlayerInfo(nil),
		},
		{
			name: "single game with win",
			games: []types.StoredGame{
				{
					Index:  "1",
					Winner: "b",
					Black:  "alice",
					Red:    "bob",
				},
			},
			expected: []types.PlayerInfo{
				{
					Index:     "alice",
					WonCount:  1,
					LostCount: 0,
				},
				{
					Index:     "bob",
					WonCount:  0,
					LostCount: 1,
				},
			},
		},
		{
			name: "two games with win",
			games: []types.StoredGame{
				{
					Index:  "1",
					Winner: "b",
					Black:  "alice",
					Red:    "bob",
				},
				{
					Index:  "2",
					Winner: "b",
					Black:  "alice",
					Red:    "carol",
				},
			},
			expected: []types.PlayerInfo{
				{
					Index:     "alice",
					WonCount:  2,
					LostCount: 0,
				},
				{
					Index:     "bob",
					WonCount:  0,
					LostCount: 1,
				},
				{
					Index:     "carol",
					WonCount:  0,
					LostCount: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		for chunk := uint64(1); chunk < 5; chunk++ {
			t.Run(fmt.Sprintf("%s chunk %d", tt.name, chunk), func(t *testing.T) {
				keeper, context := setupKeeperForV1ToV1_1Migration(t)
				ctx := sdk.UnwrapSDKContext(context)

				for _, game := range tt.games {
					keeper.SetStoredGame(ctx, game)
				}
				err := cv3keeper.MapStoredGamesReduceToPlayerInfo(ctx, keeper, chunk)
				require.Nil(t, err)

				playerInfos := keeper.GetAllPlayerInfo(ctx)
				require.Equal(t, len(tt.expected), len(playerInfos))
				require.EqualValues(t, tt.expected, playerInfos)
			})
		}
	}
}

func TestBuild10kPlayerInfosInPlace(t *testing.T) {
	chunks := []uint64{1, 10, 100, 1_000, 10_000, 100_000, 1_000_000}
	for _, chunk := range chunks {
		keeper, context := setupKeeperForV1ToV1_1Migration(t)
		ctx := sdk.UnwrapSDKContext(context)
		for id := uint64(1); id <= 100_000; id++ {
			keeper.SetStoredGame(ctx, types.StoredGame{
				Index:  strconv.FormatUint(id, 10),
				Black:  "alice",
				Red:    "bob",
				Winner: "b",
			})
		}
		before := time.Now()
		cv3keeper.MapStoredGamesReduceToPlayerInfo(ctx, keeper, chunk)
		after := time.Now()
		playerInfos := keeper.GetAllPlayerInfo(ctx)
		require.Equal(t, 2, len(playerInfos))
		require.EqualValues(t, []types.PlayerInfo{
			{
				Index:    "alice",
				WonCount: 100_000,
			},
			{
				Index:     "bob",
				LostCount: 100_000,
			},
		}, playerInfos)
		t.Logf("Chunk %d, duration %d millisec", chunk, after.Sub(before).Milliseconds())
	}
}
