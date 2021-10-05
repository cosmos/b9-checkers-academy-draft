package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidCreator = sdkerrors.Register(ModuleName, 1100, "creator address is invalid")
	ErrInvalidRed     = sdkerrors.Register(ModuleName, 1101, "red address is invalid")
	ErrInvalidBlack   = sdkerrors.Register(ModuleName, 1102, "black address is invalid")
	// this line is used by starport scaffolding # ibc/errors
)
