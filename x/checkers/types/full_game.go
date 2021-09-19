package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/rules"
)

type FullGame struct {
	Creator sdk.AccAddress
	Index   string
	Game    rules.Game
	Red     sdk.AccAddress
	Black   sdk.AccAddress
}

func (fullGame *FullGame) ToStoredGame() (storedGame *StoredGame) {
	storedGame.Creator = fullGame.Creator.String()
	storedGame.Index = fullGame.Index
	storedGame.Game = fullGame.Game.String()
	storedGame.Turn = fullGame.Game.Turn.Color
	storedGame.Red = fullGame.Red.String()
	storedGame.Black = fullGame.Black.String()
	return storedGame
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
	if err != nil {
		panic(err)
	}
	return fullGame
}
