package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidCreator          = sdkerrors.Register(ModuleName, 1100, "creator address is invalid: %s")
	ErrInvalidRed              = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrInvalidBlack            = sdkerrors.Register(ModuleName, 1102, "black address is invalid: %s")
	ErrGameNotParseable        = sdkerrors.Register(ModuleName, 1103, "game cannot be parsed")
	ErrGameNotFound            = sdkerrors.Register(ModuleName, 1104, "game by id not found: %s")
	ErrCreatorNotPlayer        = sdkerrors.Register(ModuleName, 1105, "message creator is not a player: %s")
	ErrNotPlayerTurn           = sdkerrors.Register(ModuleName, 1106, "player tried to play out of turn: %s")
	ErrWrongMove               = sdkerrors.Register(ModuleName, 1107, "wrong move")
	ErrRedAlreadyPlayed        = sdkerrors.Register(ModuleName, 1108, "red player has already played")
	ErrBlackAlreadyPlayed      = sdkerrors.Register(ModuleName, 1109, "black player has already played")
	ErrInvalidDeadline         = sdkerrors.Register(ModuleName, 1110, "deadline cannot be parsed: %s")
	ErrGameFinished            = sdkerrors.Register(ModuleName, 1111, "game is already finished")
	ErrRedCannotPay            = sdkerrors.Register(ModuleName, 1112, "red cannot pay the wager")
	ErrBlackCannotPay          = sdkerrors.Register(ModuleName, 1113, "black cannot pay the wager")
	ErrCannotFindWinnerByColor = sdkerrors.Register(ModuleName, 1114, "cannot find winner by color: %s")
	ErrNothingToPay            = sdkerrors.Register(ModuleName, 1115, "there is nothing to pay, should not have been called")
	ErrCannotRefundWager       = sdkerrors.Register(ModuleName, 1116, "cannot refund wager to: %s")
	ErrCannotPayWinnings       = sdkerrors.Register(ModuleName, 1117, "cannot pay winnings to winner")
	ErrNotInRefundState        = sdkerrors.Register(ModuleName, 1118, "game is not in a state to refund, move count: %d")
	ErrWinnerNotParseable      = sdkerrors.Register(ModuleName, 1119, "winner is not parseable: %s")
	ErrThereIsNoWinner         = sdkerrors.Register(ModuleName, 1120, "there is no winner")
	ErrInvalidDateAdded        = sdkerrors.Register(ModuleName, 1121, "dateAdded cannot be parsed: %s")
	ErrCannotAddToLeaderboard  = sdkerrors.Register(ModuleName, 1122, "cannot add to leaderboard: %s")
	// this line is used by starport scaffolding # ibc/errors
)
