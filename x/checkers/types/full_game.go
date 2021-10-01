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

func (fullGame FullGame) ToStoredGame() StoredGame {
	return StoredGame{
		Creator: fullGame.Creator.String(),
		Index:   fullGame.Index,
		Game:    fullGame.Game.String(),
		Turn:    fullGame.Game.Turn.Color,
		Red:     fullGame.Red.String(),
		Black:   fullGame.Black.String(),
	}
}

func (storedGame StoredGame) ToFullGame() (fullGame FullGame) {
	creator, err := sdk.AccAddressFromBech32(storedGame.Creator)
	game, err := rules.Parse(storedGame.Game)
	game.Turn = rules.Player{
		Color: storedGame.Turn,
	}
	red, err := sdk.AccAddressFromBech32(storedGame.Red)
	black, err := sdk.AccAddressFromBech32(storedGame.Black)
	if err != nil {
		panic(err)
	}
	return FullGame{
		Creator: creator,
		Index:   storedGame.Index,
		Game:    *game,
		Red:     red,
		Black:   black,
	}
}
