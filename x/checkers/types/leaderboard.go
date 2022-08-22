package types

import "fmt"

func (leaderboard Leaderboard) Validate() error {
	// Check for duplicated player address in winners
	winnerInfoIndexMap := make(map[string]struct{})

	for _, elem := range leaderboard.Winners {
		index := string(PlayerInfoKey(elem.PlayerAddress))
		if _, ok := winnerInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated playerAddress for winner")
		}
		winnerInfoIndexMap[index] = struct{}{}
	}
	return nil
}
