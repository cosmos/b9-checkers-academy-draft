package v1tov2_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/b9lab/checkers/x/checkers/migrations/v1tov2"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestBuildLeaderboardInPlace(t *testing.T) {
	tests := []struct {
		name        string
		playerInfos []types.PlayerInfo
		expected    []types.WinningPlayer
	}{
		{
			name:        "nothing to assemble",
			playerInfos: []types.PlayerInfo{},
			expected:    []types.WinningPlayer(nil),
		},
		{
			name: "single player no win",
			playerInfos: []types.PlayerInfo{
				{
					Index:          "alice",
					WonCount:       0,
					LostCount:      1,
					ForfeitedCount: 0,
				},
			},
			expected: []types.WinningPlayer(nil),
		},
		{
			name: "single player with win",
			playerInfos: []types.PlayerInfo{
				{
					Index:          "alice",
					WonCount:       2,
					LostCount:      1,
					ForfeitedCount: 0,
				},
			},
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "0001-01-01 00:00:00 +0000 UTC",
				},
			},
		},
		{
			name: "two players with win",
			playerInfos: []types.PlayerInfo{
				{
					Index:          "alice",
					WonCount:       2,
					LostCount:      1,
					ForfeitedCount: 0,
				},
				{
					Index:          "bob",
					WonCount:       4,
					LostCount:      0,
					ForfeitedCount: 0,
				},
			},
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "bob",
					WonCount:      4,
					DateAdded:     "0001-01-01 00:00:00 +0000 UTC",
				},
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "0001-01-01 00:00:00 +0000 UTC",
				},
			},
		},
	}
	for _, tt := range tests {
		for _, chunk := range []uint64{1, 10, v1tov2.PlayerInfoChunkSize} {
			t.Run(fmt.Sprintf("%s chunk %d", tt.name, chunk), func(t *testing.T) {
				keeper, context := setupKeeperForV1ToV2Migration(t)
				ctx := sdk.UnwrapSDKContext(context)

				for _, playerInfo := range tt.playerInfos {
					keeper.SetPlayerInfo(ctx, playerInfo)
				}
				v1tov2.MapPlayerInfosReduceToLeaderboard(ctx, keeper, chunk)

				leaderboard, found := keeper.GetLeaderboard(ctx)
				require.True(t, found)
				require.Equal(t, len(tt.expected), len(leaderboard.Winners))
				require.EqualValues(t, tt.expected, leaderboard.Winners)
			})
		}
	}
}

func Test101Players(t *testing.T) {
	for _, chunk := range []uint64{1, 10, v1tov2.PlayerInfoChunkSize} {
		t.Run(fmt.Sprintf("chunk %d", chunk), func(t *testing.T) {
			keeper, context := setupKeeperForV1ToV2Migration(t)
			ctx := sdk.UnwrapSDKContext(context)
			expectedWinners := make([]types.WinningPlayer, types.LeaderboardWinnerLength)
			for i := uint64(0); i <= types.LeaderboardWinnerLength; i++ {
				keeper.SetPlayerInfo(ctx, types.PlayerInfo{
					Index:    strconv.FormatUint(i, 10),
					WonCount: i,
				})
				if i > 0 {
					expectedWinners[types.LeaderboardWinnerLength-i] = types.WinningPlayer{
						PlayerAddress: strconv.FormatUint(i, 10),
						WonCount:      i,
						DateAdded:     "0001-01-01 00:00:00 +0000 UTC",
					}
				}
			}
			v1tov2.MapPlayerInfosReduceToLeaderboard(ctx, keeper, chunk)
			leaderboard, found := keeper.GetLeaderboard(ctx)
			require.True(t, found)
			require.Equal(t, types.LeaderboardWinnerLength, uint64(len(leaderboard.Winners)))
			require.EqualValues(t, expectedWinners, leaderboard.Winners)
		})
	}
}

func Test201Players(t *testing.T) {
	for _, chunk := range []uint64{1, 10, v1tov2.PlayerInfoChunkSize} {
		t.Run(fmt.Sprintf("chunk %d", chunk), func(t *testing.T) {

			keeper, context := setupKeeperForV1ToV2Migration(t)
			ctx := sdk.UnwrapSDKContext(context)
			expectedWinners := make([]types.WinningPlayer, types.LeaderboardWinnerLength)
			for i := uint64(0); i <= types.LeaderboardWinnerLength*2; i++ {
				keeper.SetPlayerInfo(ctx, types.PlayerInfo{
					Index:    strconv.FormatUint(i, 10),
					WonCount: i,
				})
				if i > 100 {
					expectedWinners[types.LeaderboardWinnerLength*2-i] = types.WinningPlayer{
						PlayerAddress: strconv.FormatUint(i, 10),
						WonCount:      i,
						DateAdded:     "0001-01-01 00:00:00 +0000 UTC",
					}
				}
			}
			v1tov2.MapPlayerInfosReduceToLeaderboard(ctx, keeper, chunk)
			leaderboard, found := keeper.GetLeaderboard(ctx)
			require.True(t, found)
			require.Equal(t, types.LeaderboardWinnerLength, uint64(len(leaderboard.Winners)))
			require.EqualValues(t, expectedWinners, leaderboard.Winners)
		})
	}
}
