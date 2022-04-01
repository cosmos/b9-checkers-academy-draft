package types

import (
	"errors"
	"fmt"
	"time"

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
	if errGame != nil {
		return nil, sdkerrors.Wrapf(errGame, ErrGameNotParseable.Error())
	}
	game.Turn = rules.StringPieces[storedGame.Turn].Player
	if game.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), ErrGameNotParseable.Error())
	}
	return game, nil
}

func (storedGame *StoredGame) GetDeadlineAsTime() (deadline time.Time, err error) {
	deadline, errDeadline := time.Parse(DeadlineLayout, storedGame.Deadline)
	return deadline, sdkerrors.Wrapf(errDeadline, ErrInvalidDeadline.Error(), storedGame.Deadline)
}

func GetNextDeadline(ctx sdk.Context) time.Time {
	return ctx.BlockTime().Add(MaxTurnDurationInSeconds)
}

func FormatDeadline(deadline time.Time) string {
	return deadline.UTC().Format(DeadlineLayout)
}

func (storedGame *StoredGame) GetPlayerAddress(color string) (address sdk.AccAddress, found bool, err error) {
	red, err := storedGame.GetRedAddress()
	if err != nil {
		return nil, false, err
	}
	black, err := storedGame.GetBlackAddress()
	if err != nil {
		return nil, false, err
	}
	address, found = map[string]sdk.AccAddress{
		rules.RED_PLAYER.Color:   red,
		rules.BLACK_PLAYER.Color: black,
	}[color]
	return address, found, nil
}

func (storedGame *StoredGame) GetWinnerAddress() (address sdk.AccAddress, found bool, err error) {
	address, found, err = storedGame.GetPlayerAddress(storedGame.Winner)
	return address, found, err
}

func (storedGame *StoredGame) GetWagerCoin() (wager sdk.Coin) {
	return sdk.NewCoin(storedGame.Token, sdk.NewInt(int64(storedGame.Wager)))
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetCreatorAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetBlackAddress()
	return err
}
