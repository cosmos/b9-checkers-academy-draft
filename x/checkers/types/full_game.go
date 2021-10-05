package types

import (
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
	}
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
	return FullGame{
		Creator:   creator,
		Index:     storedGame.Index,
		Game:      *game,
		Red:       red,
		Black:     black,
		MoveCount: storedGame.MoveCount,
	}
}
