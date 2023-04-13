package types_test

import (
	"testing"

	"github.com/b9lab/checkers/x/leaderboard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				Leaderboard: types.Leaderboard{
					Winners: []types.Winner{},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated winnerPlayer",
			genState: &types.GenesisState{
				Leaderboard: types.Leaderboard{
					Winners: []types.Winner{
						{
							Address: "cosmos123",
						},
						{
							Address: "cosmos123",
						},
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDefaultGenesisState_ExpectedInitial(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			Leaderboard: types.Leaderboard{
				Winners: []types.Winner{},
			},
			Params: types.Params{
				Length: 100,
			},
		},
		types.DefaultGenesis())
}
