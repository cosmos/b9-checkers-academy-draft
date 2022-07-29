package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidBlack            = sdkerrors.Register(ModuleName, 1100, "black address is invalid: %s")
	ErrInvalidRed              = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrGameNotParseable        = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
	ErrGameNotFound            = sdkerrors.Register(ModuleName, 1103, "game by id not found")
	ErrCreatorNotPlayer        = sdkerrors.Register(ModuleName, 1104, "message creator is not a player")
	ErrNotPlayerTurn           = sdkerrors.Register(ModuleName, 1105, "player tried to play out of turn")
	ErrWrongMove               = sdkerrors.Register(ModuleName, 1106, "wrong move")
	ErrBlackAlreadyPlayed      = sdkerrors.Register(ModuleName, 1107, "black player has already played")
	ErrRedAlreadyPlayed        = sdkerrors.Register(ModuleName, 1108, "red player has already played")
	ErrInvalidDeadline         = sdkerrors.Register(ModuleName, 1109, "deadline cannot be parsed: %s")
	ErrGameFinished            = sdkerrors.Register(ModuleName, 1110, "game is already finished")
	ErrCannotFindWinnerByColor = sdkerrors.Register(ModuleName, 1111, "cannot find winner by color: %s")
)
