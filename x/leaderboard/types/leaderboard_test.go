package types_test

import (
	"testing"
	"time"

	"github.com/b9lab/checkers/x/leaderboard/testutil"
	"github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func TestCandidateGetWinnerAtTime(t *testing.T) {
	now := time.Now()
	timestamp := now.Unix()
	aliceAddress, err := sdk.AccAddressFromBech32(alice)
	require.Nil(t, err)
	candidate := types.Candidate{
		Address:  aliceAddress,
		WonCount: 23,
	}
	winner := candidate.GetWinnerAtTime(now)
	require.EqualValues(t, types.Winner{
		Address:  alice,
		WonCount: 23,
		AddedAt:  uint64(timestamp),
	}, winner)
}

func TestSortWinners(t *testing.T) {
	tests := []struct {
		name     string
		unsorted []types.Winner
		sorted   []types.Winner
	}{
		{
			name:     "sort empty",
			unsorted: []types.Winner{},
			sorted:   []types.Winner{},
		},
		{
			name: "sort unique",
			unsorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
		},
		{
			name: "sort already two sorted",
			unsorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
				{
					Address:  bob,
					WonCount: 1,
					AddedAt:  1001,
				},
			},
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
				{
					Address:  bob,
					WonCount: 1,
					AddedAt:  1001,
				},
			},
		},
		{
			name: "sort two not sorted",
			unsorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 1,
					AddedAt:  1001,
				},
				{
					Address:  bob,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
			sorted: []types.Winner{
				{
					Address:  bob,
					WonCount: 2,
					AddedAt:  1000,
				},
				{
					Address:  alice,
					WonCount: 1,
					AddedAt:  1001,
				},
			},
		},
		{
			name: "sort two not sorted by date",
			unsorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
				{
					Address:  bob,
					WonCount: 2,
					AddedAt:  1001,
				},
			},
			sorted: []types.Winner{
				{
					Address:  bob,
					WonCount: 2,
					AddedAt:  1001,
				},
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leaderboard := types.Leaderboard{
				Winners: tt.unsorted,
			}
			leaderboard.SortWinners()
			sorted := leaderboard.Winners
			require.Equal(t, len(tt.sorted), len(sorted))
			require.EqualValues(t, tt.sorted, sorted)
		})
	}
}

func TestAddCandidatesAtNow(t *testing.T) {
	aliceAdd, err := sdk.AccAddressFromBech32(alice)
	require.Nil(t, err)
	bobAdd, err := sdk.AccAddressFromBech32(bob)
	require.Nil(t, err)
	tests := []struct {
		name       string
		sorted     []types.Winner
		candidates []types.Candidate
		now        int64
		expected   []types.Winner
	}{
		{
			name:   "add to empty",
			sorted: []types.Winner{},
			candidates: []types.Candidate{{
				Address:  aliceAdd,
				WonCount: 2,
			}},
			now: 1000,
			expected: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
		},
		{
			name: "update alice alone",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 1,
					AddedAt:  999,
				},
			},
			candidates: []types.Candidate{{
				Address:  aliceAdd,
				WonCount: 2,
			}},
			now: 1000,
			expected: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
		},
		{
			name: "bob added ahead by count",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
			candidates: []types.Candidate{{
				Address:  bobAdd,
				WonCount: 3,
			}},
			now: 999,
			expected: []types.Winner{
				{
					Address:  bob,
					WonCount: 3,
					AddedAt:  999,
				},
				{
					Address:  alice,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
		},
		{
			name: "bob added ahead by time",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  999,
				},
			},
			candidates: []types.Candidate{{
				Address:  bobAdd,
				WonCount: 3,
			}},
			now: 1000,
			expected: []types.Winner{
				{
					Address:  bob,
					WonCount: 3,
					AddedAt:  1000,
				},
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  999,
				},
			},
		},
		{
			name: "bob added behind by count",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  999,
				},
			},
			candidates: []types.Candidate{{
				Address:  bobAdd,
				WonCount: 2,
			}},
			now: 1000,
			expected: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  999,
				},
				{
					Address:  bob,
					WonCount: 2,
					AddedAt:  1000,
				},
			},
		},
		{
			name: "bob added behind by time",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
			},
			candidates: []types.Candidate{{
				Address:  bobAdd,
				WonCount: 3,
			}},
			now: 999,
			expected: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
				{
					Address:  bob,
					WonCount: 3,
					AddedAt:  999,
				},
			},
		},
		{
			name: "alice unchanged by more ancient time",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
			},
			candidates: []types.Candidate{{
				Address:  aliceAdd,
				WonCount: 3,
			}},
			now: 999,
			expected: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
			},
		},
		{
			name: "alice unchanged by more recent time",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
			},
			candidates: []types.Candidate{{
				Address:  aliceAdd,
				WonCount: 3,
			}},
			now: 1001,
			expected: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
			},
		},
		{
			name: "update alice ahead",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
				{
					Address:  bob,
					WonCount: 3,
					AddedAt:  1001,
				},
			},
			candidates: []types.Candidate{{
				Address:  aliceAdd,
				WonCount: 5,
			}},
			now: 1000,
			expected: []types.Winner{
				{
					Address:  alice,
					WonCount: 5,
					AddedAt:  1000,
				},
				{
					Address:  bob,
					WonCount: 3,
					AddedAt:  1001,
				},
			},
		},
		{
			name: "update bob behind",
			sorted: []types.Winner{
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
				{
					Address:  bob,
					WonCount: 3,
					AddedAt:  1000,
				},
			},
			candidates: []types.Candidate{{
				Address:  bobAdd,
				WonCount: 4,
			}},
			now: 1000,
			expected: []types.Winner{
				{
					Address:  bob,
					WonCount: 4,
					AddedAt:  1000,
				},
				{
					Address:  alice,
					WonCount: 3,
					AddedAt:  1000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := types.AddCandidatesAtNow(tt.sorted, time.Unix(tt.now, 0), tt.candidates)
			require.Equal(t, len(tt.expected), len(actual))
			require.EqualValues(t, tt.expected, actual)
			require.NoError(t, types.Leaderboard{Winners: actual}.Validate())
		})
	}
}
