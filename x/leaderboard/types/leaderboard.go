package types

import (
	fmt "fmt"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (leaderboard Leaderboard) Validate() error {
	// Check for duplicated player address in winners
	winnerInfoIndexMap := make(map[string]struct{})

	for index, elem := range leaderboard.Winners {
		if _, ok := winnerInfoIndexMap[elem.Address]; ok {
			return fmt.Errorf("duplicated address %s at index %d", elem.Address, index)
		}
		winnerInfoIndexMap[elem.Address] = struct{}{}
	}
	return nil
}

func (candidate Candidate) GetAccAddress() string {
	return sdk.AccAddress(candidate.Address).String()
}

func (candidate Candidate) GetWinnerAtTime(now time.Time) Winner {
	return Winner{
		Address:  candidate.GetAccAddress(),
		WonCount: candidate.WonCount,
		AddedAt:  uint64(now.Unix()),
	}
}

func SortWinners(winners []Winner) {
	sort.SliceStable(winners[:], func(i, j int) bool {
		if winners[i].WonCount > winners[j].WonCount {
			return true
		}
		if winners[i].WonCount < winners[j].WonCount {
			return false
		}
		return winners[i].AddedAt > winners[j].AddedAt
	})
}

func (leaderboard Leaderboard) SortWinners() {
	SortWinners(leaderboard.Winners)
}

func MapWinners(winners []Winner, length int) map[string]Winner {
	mapped := make(map[string]Winner, length)
	for _, winner := range winners {
		already, found := mapped[winner.Address]
		if !found {
			mapped[winner.Address] = winner
		} else if already.WonCount < winner.WonCount {
			mapped[winner.Address] = winner
		}
	}
	return mapped
}

func AddCandidatesAtNow(winners []Winner, now time.Time, candidates []Candidate) (updated []Winner) {
	mapped := MapWinners(winners, len(winners)+len(candidates))
	for _, candidate := range candidates {
		if candidate.WonCount < 1 {
			continue
		}
		candidateWinner := candidate.GetWinnerAtTime(now)
		already, found := mapped[candidateWinner.Address]
		if !found {
			mapped[candidateWinner.Address] = candidateWinner
		} else if already.WonCount < candidateWinner.WonCount {
			mapped[candidateWinner.Address] = candidateWinner
		}
	}
	updated = make([]Winner, 0, len(mapped))
	for _, winner := range mapped {
		updated = append(updated, winner)
	}
	SortWinners(updated)
	return updated
}
