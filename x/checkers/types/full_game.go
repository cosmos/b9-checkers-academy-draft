package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/xavierlepretre/checkers/x/checkers/rules"
)

func (storedGame *StoredGame) GetCreatorAddress() (creator sdk.AccAddress, err error) {
	creator, errCreator := sdk.AccAddressFromBech32(storedGame.Creator)
	return creator, sdkerrors.Wrapf(errCreator, ErrInvalidCreator.Error(), storedGame.Creator)
}

func (storedGame *StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame *StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame *StoredGame) ParseGame() (game *rules.Game, err error) {
	game, errGame := rules.Parse(storedGame.Game)
	if err != nil {
		return game, sdkerrors.Wrapf(errGame, ErrGameNotParseable.Error())
	}
	game.Turn = rules.Player{
		Color: storedGame.Turn,
	}
	return game, nil
}
