package types

import (
// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Leaderboard: Leaderboard{
			Winners: []Winner{},
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Validate Leaderboard
	if err := gs.Leaderboard.Validate(); err != nil {
		return err
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
