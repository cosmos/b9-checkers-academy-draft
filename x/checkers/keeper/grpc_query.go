package keeper

import (
	"github.com/b9lab/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
