package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisIsCorrect(t *testing.T) {
	require.EqualValues(t,
		&GenesisState{
			StoredGameList: []*StoredGame{},
			NextGame:       &NextGame{"", uint64(1)},
		},
		DefaultGenesis())
}
