package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidCreator     = sdkerrors.Register(ModuleName, 1100, "creator address is invalid")
	ErrInvalidRed         = sdkerrors.Register(ModuleName, 1101, "red address is invalid")
	ErrInvalidBlack       = sdkerrors.Register(ModuleName, 1102, "black address is invalid")
	ErrGameNotFound       = sdkerrors.Register(ModuleName, 1103, "game by id not found")
	ErrCreatorNotPlayer   = sdkerrors.Register(ModuleName, 1104, "message creator is not a player")
	ErrNotPlayerTurn      = sdkerrors.Register(ModuleName, 1105, "player tried to play out of turn")
	ErrWrongMove          = sdkerrors.Register(ModuleName, 1106, "wrong move")
	ErrRedAlreadyPlayed   = sdkerrors.Register(ModuleName, 1107, "red player has already played")
	ErrBlackAlreadyPlayed = sdkerrors.Register(ModuleName, 1108, "black player has already played")
	ErrGameFinished       = sdkerrors.Register(ModuleName, 1109, "game is already finished")
	// this line is used by starport scaffolding # ibc/errors
)
