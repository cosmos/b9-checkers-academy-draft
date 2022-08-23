package types_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestSortStringifiedWinners(t *testing.T) {
	tests := []struct {
		name     string
		unsorted []types.WinningPlayer
		sorted   []types.WinningPlayer
		err      error
	}{
		{
			name: "cannot parse date",
			unsorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "200T-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			sorted: []types.WinningPlayer{},
			err:    errors.New("dateAdded cannot be parsed: 200T-01-02 15:05:05.999999999 +0000 UTC: parsing time \"200T-01-02 15:05:05.999999999 +0000 UTC\" as \"2006-01-02 15:04:05.999999999 +0000 UTC\": cannot parse \"-01-02 15:05:05.999999999 +0000 UTC\" as \"2006\""),
		},
		{
			name:     "sort empty",
			unsorted: []types.WinningPlayer{},
			sorted:   []types.WinningPlayer{},
		},
		{
			name: "sort unique",
			unsorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "sort already two sorted",
			unsorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      1,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      1,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "sort two not sorted",
			unsorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      1,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "bob",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "alice",
					WonCount:      1,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "sort two not sorted by date",
			unsorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:04:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "bob",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:04:05.999999999 +0000 UTC",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leaderboard := types.Leaderboard{
				Winners: tt.unsorted,
			}
			parsed, err := leaderboard.ParseWinners()
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
			types.SortWinners(parsed)
			sorted := types.StringifyWinners(parsed)
			require.Equal(t, len(tt.sorted), len(sorted))
			require.EqualValues(t, tt.sorted, sorted)
		})
	}
}

func TestUpdatePlayerInfoAtNow(t *testing.T) {
	tests := []struct {
		name      string
		sorted    []types.WinningPlayer
		candidate types.PlayerInfo
		now       string
		expected  []types.WinningPlayer
	}{
		{
			name:   "add to empty",
			sorted: []types.WinningPlayer{},
			candidate: types.PlayerInfo{
				Index:    "alice",
				WonCount: 2,
			},
			now: "2006-01-02 15:05:05.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "update alice alone",
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      1,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			candidate: types.PlayerInfo{
				Index:    "alice",
				WonCount: 2,
			},
			now: "2006-01-02 15:05:05.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "bob ahead by count",
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			candidate: types.PlayerInfo{
				Index:    "bob",
				WonCount: 3,
			},
			now: "2006-01-02 15:05:05.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "bob",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "alice",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "bob ahead by time",
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			candidate: types.PlayerInfo{
				Index:    "bob",
				WonCount: 3,
			},
			now: "2006-01-02 15:05:06.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "bob",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:06.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "bob behind by count",
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			candidate: types.PlayerInfo{
				Index:    "bob",
				WonCount: 2,
			},
			now: "2006-01-02 15:05:05.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      2,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "bob behind by time",
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
			candidate: types.PlayerInfo{
				Index:    "bob",
				WonCount: 3,
			},
			now: "2006-01-02 15:05:04.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:04.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "update alice ahead",
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:04.999999999 +0000 UTC",
				},
			},
			candidate: types.PlayerInfo{
				Index:    "alice",
				WonCount: 5,
			},
			now: "2006-01-02 15:05:08.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      5,
					DateAdded:     "2006-01-02 15:05:08.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:04.999999999 +0000 UTC",
				},
			},
		},
		{
			name: "update bob behind",
			sorted: []types.WinningPlayer{
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "bob",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:04.999999999 +0000 UTC",
				},
			},
			candidate: types.PlayerInfo{
				Index:    "bob",
				WonCount: 4,
			},
			now: "2006-01-02 15:05:05.999999999 +0000 UTC",
			expected: []types.WinningPlayer{
				{
					PlayerAddress: "bob",
					WonCount:      4,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
				{
					PlayerAddress: "alice",
					WonCount:      3,
					DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now, err := types.ParseDateAddedAsTime(tt.now)
			require.NoError(t, err)
			leaderboard := types.Leaderboard{
				Winners: tt.sorted,
			}
			err = leaderboard.UpdatePlayerInfoAtNow(now, tt.candidate)
			require.NoError(t, err)
			require.Equal(t, len(tt.expected), len(leaderboard.Winners))
			require.EqualValues(t, tt.expected, leaderboard.Winners)
			require.NoError(t, leaderboard.Validate())
		})
	}
}

func makeMaxLengthSortedWinningPlayers() []types.WinningPlayer {
	sorted := make([]types.WinningPlayer, 100)
	for i := uint64(0); i < 100; i++ {
		sorted[i] = types.WinningPlayer{
			PlayerAddress: strconv.FormatUint(i, 10),
			WonCount:      101 - i,
			DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
		}
	}
	return sorted
}

func TestUpdatePlayerInfoAtNowTooLongNoAdd(t *testing.T) {
	beforeWinners := makeMaxLengthSortedWinningPlayers()
	now, err := types.ParseDateAddedAsTime("2006-01-02 15:05:05.999999999 +0000 UTC")
	require.NoError(t, err)
	leaderboard := types.Leaderboard{
		Winners: beforeWinners,
	}
	err = leaderboard.UpdatePlayerInfoAtNow(now, types.PlayerInfo{
		Index:    "100",
		WonCount: 1,
	})
	require.NoError(t, err)
	require.Equal(t, len(beforeWinners), len(leaderboard.Winners))
	require.EqualValues(t, beforeWinners, leaderboard.Winners)
	require.NoError(t, leaderboard.Validate())
}

func TestUpdatePlayerInfoAtNowLongDropSmallest(t *testing.T) {
	beforeWinners := makeMaxLengthSortedWinningPlayers()
	now, err := types.ParseDateAddedAsTime("2006-01-02 15:05:06.999999999 +0000 UTC")
	require.NoError(t, err)
	leaderboard := types.Leaderboard{
		Winners: beforeWinners,
	}
	err = leaderboard.UpdatePlayerInfoAtNow(now, types.PlayerInfo{
		Index:    "100",
		WonCount: 2,
	})
	require.NoError(t, err)
	beforeWinners[99] = types.WinningPlayer{
		PlayerAddress: "100",
		WonCount:      2,
		DateAdded:     "2006-01-02 15:05:06.999999999 +0000 UTC",
	}
	require.Equal(t, len(beforeWinners), len(leaderboard.Winners))
	require.EqualValues(t, beforeWinners, leaderboard.Winners)
	require.NoError(t, leaderboard.Validate())
}

func TestUpdatePlayerInfoAtNowLongUpdateWithin(t *testing.T) {
	beforeWinners := makeMaxLengthSortedWinningPlayers()
	now, err := types.ParseDateAddedAsTime("2006-01-02 15:05:06.999999999 +0000 UTC")
	require.NoError(t, err)
	leaderboard := types.Leaderboard{
		Winners: beforeWinners,
	}
	err = leaderboard.UpdatePlayerInfoAtNow(now, types.PlayerInfo{
		Index:    "9",
		WonCount: 93,
	})
	require.NoError(t, err)
	beforeWinners[8] = types.WinningPlayer{
		PlayerAddress: "9",
		WonCount:      93,
		DateAdded:     "2006-01-02 15:05:06.999999999 +0000 UTC",
	}
	beforeWinners[9] = types.WinningPlayer{
		PlayerAddress: "8",
		WonCount:      93,
		DateAdded:     "2006-01-02 15:05:05.999999999 +0000 UTC",
	}
	require.Equal(t, len(beforeWinners), len(leaderboard.Winners))
	require.EqualValues(t, beforeWinners, leaderboard.Winners)
	require.NoError(t, leaderboard.Validate())
}
