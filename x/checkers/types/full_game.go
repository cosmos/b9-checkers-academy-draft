package types

import (
	"strconv"
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

func (fullGame *FullGame) ToStoredGame() (storedGame *StoredGame) {
	storedGame.Creator = fullGame.Creator.String()
	storedGame.Index = fullGame.Index
	storedGame.Game = fullGame.Game.String()
	storedGame.Turn = fullGame.Game.Turn.Color
	storedGame.Red = fullGame.Red.String()
	storedGame.Black = fullGame.Black.String()
	storedGame.MoveCount = strconv.FormatUint(fullGame.MoveCount, 10)
	storedGame.BeforeId = fullGame.BeforeId
	storedGame.AfterId = fullGame.BeforeId
	storedGame.Deadline = fullGame.Deadline.UTC().Format(DeadlineLayout)
	storedGame.Winner = fullGame.Winner
	storedGame.Wager = fullGame.Wager.Amount.Uint64()
	storedGame.Token = fullGame.Wager.Denom
	return storedGame
}

func (storedGame *StoredGame) GetDeadlineAsTime() (deadline time.Time, err error) {
	deadline, err = time.Parse(DeadlineLayout, storedGame.Deadline)
	return deadline, err
}

func (storedGame *StoredGame) ToFullGame() (fullGame *FullGame) {
	var err error
	fullGame.Creator, err = sdk.AccAddressFromBech32(storedGame.Creator)
	fullGame.Index = storedGame.Index
	var game *rules.Game
	game, err = rules.Parse(storedGame.Game)
	game.Turn = rules.Player{
		Color: storedGame.Turn,
	}
	fullGame.Game = *game
	fullGame.Red, err = sdk.AccAddressFromBech32(storedGame.Red)
	fullGame.Black, err = sdk.AccAddressFromBech32(storedGame.Black)
	fullGame.MoveCount, err = strconv.ParseUint(storedGame.MoveCount, 10, 64)
	fullGame.BeforeId = storedGame.BeforeId
	fullGame.AfterId = storedGame.AfterId
	fullGame.Deadline, err = storedGame.GetDeadlineAsTime()
	fullGame.Winner = storedGame.Winner
	fullGame.Wager = sdk.NewCoin(storedGame.Token, sdk.NewInt(int64(storedGame.Wager)))
	if err != nil {
		panic(err)
	}
	return fullGame
}
