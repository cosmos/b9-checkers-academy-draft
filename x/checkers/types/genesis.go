package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		Leaderboard: &Leaderboard{
			Winners: []*WinningPlayer{},
		},
		PlayerInfoList: []*PlayerInfo{},
		StoredGameList: []*StoredGame{},
		NextGame: &NextGame{
			Creator:  "",
			IdValue:  uint64(0),
			FifoHead: NoFifoIdKey,
			FifoTail: NoFifoIdKey,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in playerInfo
	playerInfoIndexMap := make(map[string]bool)

	for _, elem := range gs.PlayerInfoList {
		if _, ok := playerInfoIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for playerInfo")
		}
		playerInfoIndexMap[elem.Index] = true
	}
	// Check for duplicated index in storedGame
	storedGameIndexMap := make(map[string]bool)

	for _, elem := range gs.StoredGameList {
		if _, ok := storedGameIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for storedGame")
		}
		storedGameIndexMap[elem.Index] = true
	}

	return nil
}
