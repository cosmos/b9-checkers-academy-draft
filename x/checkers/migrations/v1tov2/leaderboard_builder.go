package v1tov2

import (
	"time"

	"github.com/xavierlepretre/checkers/x/checkers/types"
)

const (
	// Adjust this length to obtain the best performance over a large map.
	IntermediaryPlayerLength = types.LeaderboardWinnerLength * 2
)

// Make a leaderboard that stays in memory during the genesis conversion, and that has the required capacity
// as per the PopulateLeaderboardWith algorithm.
func CreateLeaderboardForGenesis() *types.Leaderboard {
	return &types.Leaderboard{
		Winners: make([]*types.WinningPlayer, 0, types.LeaderboardWinnerLength+IntermediaryPlayerLength),
	}
}

func PopulateLeaderboardWith(leaderboard *types.Leaderboard, additionalPlayers *map[string]*types.PlayerInfo, now time.Time) (err error) {
	partialPlayers := make([]*types.PlayerInfo, IntermediaryPlayerLength)
	for _, playerInfo := range *additionalPlayers {
		partialPlayers = append(partialPlayers, playerInfo)
		if len(partialPlayers) >= cap(partialPlayers) {
			leaderboard.AddCandidatesAndSortAtNow(now, partialPlayers)
			partialPlayers = partialPlayers[:0]
		}
	}
	leaderboard.AddCandidatesAndSortAtNow(now, partialPlayers)
	return nil
}
