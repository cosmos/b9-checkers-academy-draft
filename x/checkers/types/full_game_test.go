package types

import (
	"strings"
	"testing"

	"github.com/b9lab/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() *StoredGame {
	return &StoredGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Index:   "1",
		Game:    rules.New().String(),
		Turn:    "b",
	}
}

func TestCanGetAddressCreator(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	creator, err2 := GetStoredGame1().GetCreatorAddress()
	require.Equal(t, aliceAddress, creator)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongCreator(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Creator = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"
	creator, err := storedGame.GetCreatorAddress()
	require.Nil(t, creator)
	require.EqualError(t,
		err,
		"creator address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: checksum failed. Expected 3xn9d3, got 3xn9d4.")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetAddressBlack(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	black, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, bobAddress, black)
	require.Nil(t, err2)
	require.Nil(t, err1)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h"
	black, err := storedGame.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t,
		err,
		"black address is invalid: cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h: decoding bech32 failed: checksum failed. Expected xqhc8g, got xqhc8h.")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetAddressRed(t *testing.T) {
	carolAddress, err1 := sdk.AccAddressFromBech32(carol)
	red, err2 := GetStoredGame1().GetRedAddress()
	require.Equal(t, carolAddress, red)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongRed(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd8"
	red, err := storedGame.GetRedAddress()
	require.Nil(t, red)
	require.EqualError(t,
		err,
		"red address is invalid: cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd8: decoding bech32 failed: checksum failed. Expected 4yuxd7, got 4yuxd8.")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameCorrect(t *testing.T) {
	game, err := GetStoredGame1().ParseGame()
	require.EqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}

func TestParseGameCanIfChangedOk(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Game = strings.Replace(storedGame.Game, "b", "r", 1)
	game, err := storedGame.ParseGame()
	require.NotEqualValues(t, rules.New().Pieces, game)
	require.Nil(t, err)
}

func TestParseGameWrongPieceColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Game = strings.Replace(storedGame.Game, "b", "w", 1)
	game, err := storedGame.ParseGame()
	require.Nil(t, game)
	require.EqualError(t, err, "game cannot be parsed: invalid board, invalid piece at 1, 0")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameWrongTurnColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Turn = "w"
	game, err := storedGame.ParseGame()
	require.Nil(t, game)
	require.EqualError(t, err, "game cannot be parsed: Turn: w")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestGameValidateOk(t *testing.T) {
	storedGame := GetStoredGame1()
	require.NoError(t, storedGame.Validate())
}
