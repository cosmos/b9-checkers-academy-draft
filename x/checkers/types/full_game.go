package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/rules"
)

type FullGame struct {
	Creator   sdk.AccAddress
	Index     string
	Game      rules.Game
	Red       sdk.AccAddress
	Black     sdk.AccAddress
	MoveCount uint64
	BeforeId  string
	AfterId   string
	Deadline  time.Time
	Winner    string
	Wager     sdk.Coin
}

func (fullgame *FullGame) GetPlayerAddress(color string) (address sdk.AccAddress, found bool) {
	address, found = map[string]sdk.AccAddress{
		rules.RED_PLAYER.Color:   fullgame.Red,
		rules.BLACK_PLAYER.Color: fullgame.Black,
	}[color]
	return address, found
}

func (fullGame *FullGame) GetWinnerAddress() (address sdk.AccAddress, found bool) {
	return fullGame.GetPlayerAddress(fullGame.Winner)
}

func (fullGame FullGame) ToStoredGame() StoredGame {
	return StoredGame{
		Creator:   fullGame.Creator.String(),
		Index:     fullGame.Index,
		Game:      fullGame.Game.String(),
		Turn:      fullGame.Game.Turn.Color,
		Red:       fullGame.Red.String(),
		Black:     fullGame.Black.String(),
		MoveCount: fullGame.MoveCount,
		BeforeId:  fullGame.BeforeId,
		AfterId:   fullGame.BeforeId,
		Deadline:  fullGame.Deadline.UTC().Format(DeadlineLayout),
		Winner:    fullGame.Winner,
		Wager:     fullGame.Wager.Amount.Uint64(),
	}
}

func (storedGame *StoredGame) GetDeadlineAsTime() (deadline time.Time, err error) {
	deadline, err = time.Parse(DeadlineLayout, storedGame.Deadline)
	return deadline, err
}

func (storedGame StoredGame) ToFullGame() (fullGame FullGame) {
	creator, err := sdk.AccAddressFromBech32(storedGame.Creator)
	if err != nil {
		panic(err)
	}
	game, err := rules.Parse(storedGame.Game)
	if err != nil {
		panic(err)
	}
	game.Turn = rules.Player{
		Color: storedGame.Turn,
	}
	red, err := sdk.AccAddressFromBech32(storedGame.Red)
	if err != nil {
		panic(err)
	}
	black, err := sdk.AccAddressFromBech32(storedGame.Black)
	if err != nil {
		panic(err)
	}
	deadline, err := storedGame.GetDeadlineAsTime()
	if err != nil {
		panic(err)
	}
	return FullGame{
		Creator:   creator,
		Index:     storedGame.Index,
		Game:      *game,
		Red:       red,
		Black:     black,
		MoveCount: storedGame.MoveCount,
		BeforeId:  storedGame.BeforeId,
		AfterId:   storedGame.AfterId,
		Deadline:  deadline,
		Winner:    storedGame.Winner,
		Wager:     sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(storedGame.Wager))),
	}
}
