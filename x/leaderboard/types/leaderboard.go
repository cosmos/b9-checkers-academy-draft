package types

import (
	fmt "fmt"
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
