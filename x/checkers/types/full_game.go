package types

import (
	"fmt"

	"github.com/b9lab/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}
	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
	}
	return board, nil
}

func (storedGame StoredGame) GetPlayerAddress(color string) (address sdk.AccAddress, found bool, err error) {
	black, err := storedGame.GetBlackAddress()
	if err != nil {
		return nil, false, err
	}
	red, err := storedGame.GetRedAddress()
	if err != nil {
		return nil, false, err
	}
	address, found = map[string]sdk.AccAddress{
		rules.PieceStrings[rules.BLACK_PLAYER]: black,
		rules.PieceStrings[rules.RED_PLAYER]:   red,
	}[color]
	return address, found, nil
}

func (storedGame StoredGame) GetWinnerAddress() (address sdk.AccAddress, found bool, err error) {
	return storedGame.GetPlayerAddress(storedGame.Winner)
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	return err
}
