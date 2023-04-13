package keeper

import (
	"github.com/b9lab/checkers/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
