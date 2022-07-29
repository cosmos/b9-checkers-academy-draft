package types_test

import (
	"testing"

	"github.com/b9lab/checkers/x/checkers/types"
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

				SystemInfo: types.SystemInfo{
					NextId: 39,
				},
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated storedGame",
			genState: &types.GenesisState{
				StoredGameList: []types.StoredGame{
					{
						Index: "0",
					},
					{
						Index: "0",
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

func TestDefaultGenesisState_ExpectedInitialNextId(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			StoredGameList: []types.StoredGame{},
			SystemInfo: types.SystemInfo{
				NextId:        uint64(1),
				FifoHeadIndex: "-1",
				FifoTailIndex: "-1",
			},
		},
		types.DefaultGenesis())
}
