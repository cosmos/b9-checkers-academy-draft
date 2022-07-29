package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidBlack         = sdkerrors.Register(ModuleName, 1100, "black address is invalid: %s")
	ErrInvalidRed           = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrGameNotParseable     = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
	ErrInvalidGameIndex     = sdkerrors.Register(ModuleName, 1103, "game index is invalid")
	ErrInvalidPositionIndex = sdkerrors.Register(ModuleName, 1104, "position index is invalid")
	ErrMoveAbsent           = sdkerrors.Register(ModuleName, 1105, "there is no move")
	ErrGameNotFound         = sdkerrors.Register(ModuleName, 1106, "game by id not found")
	ErrCreatorNotPlayer     = sdkerrors.Register(ModuleName, 1107, "message creator is not a player")
	ErrNotPlayerTurn        = sdkerrors.Register(ModuleName, 1108, "player tried to play out of turn")
	ErrWrongMove            = sdkerrors.Register(ModuleName, 1109, "wrong move")
	ErrGameFinished         = sdkerrors.Register(ModuleName, 1110, "game is already finished")
)
