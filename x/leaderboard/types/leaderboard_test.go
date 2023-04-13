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
