package types_test

import (
	"strings"
	"testing"
	"time"

	"github.com/b9lab/checkers/x/checkers/rules"
	"github.com/b9lab/checkers/x/checkers/testutil"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func GetStoredGame1() types.StoredGame {
	return types.StoredGame{
		Black:       alice,
		Red:         bob,
		Index:       "1",
		Board:       rules.New().String(),
		Turn:        "b",
		MoveCount:   0,
		BeforeIndex: types.NoFifoIndex,
		AfterIndex:  types.NoFifoIndex,
		Deadline:    types.DeadlineLayout,
		Winner:      rules.PieceStrings[rules.NO_PLAYER],
	}
}

func TestCanGetAddressBlack(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	black, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, aliceAddress, black)
	require.Nil(t, err2)
	require.Nil(t, err1)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4" // Bad last digit
	black, err := storedGame.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t,
		err,
		"black address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetAddressRed(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	red, err2 := GetStoredGame1().GetRedAddress()
	require.Equal(t, bobAddress, red)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongRed(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h" // Bad last digit
	red, err := storedGame.GetRedAddress()
	require.Nil(t, red)
	require.EqualError(t,
		err,
		"red address is invalid: cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h: decoding bech32 failed: invalid checksum (expected xqhc8g got xqhc8h)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameCorrect(t *testing.T) {
	game, err := GetStoredGame1().ParseGame()
	require.EqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}

func TestParseGameCanIfChangedOk(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Board = strings.Replace(storedGame.Board, "b", "r", 1)
	game, err := storedGame.ParseGame()
	require.NotEqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}

func TestParseGameWrongPieceColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Board = strings.Replace(storedGame.Board, "b", "w", 1)
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

func TestParseDeadlineCorrect(t *testing.T) {
	deadline, err := GetStoredGame1().GetDeadlineAsTime()
	require.Nil(t, err)
	require.Equal(t, time.Time(time.Date(2006, time.January, 2, 15, 4, 5, 999999999, time.UTC)), deadline)
}

func TestParseDeadlineMissingMonth(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Deadline = "2006-02 15:04:05.999999999 +0000 UTC"
	_, err := storedGame.GetDeadlineAsTime()
	require.EqualError(t,
		err,
		"deadline cannot be parsed: 2006-02 15:04:05.999999999 +0000 UTC: parsing time \"2006-02 15:04:05.999999999 +0000 UTC\" as \"2006-01-02 15:04:05.999999999 +0000 UTC\": cannot parse \" 15:04:05.999999999 +0000 UTC\" as \"-\"")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestGetPlayerAddressBlackCorrect(t *testing.T) {
	storedGame := GetStoredGame1()
	black, found, err := storedGame.GetPlayerAddress("b")
	require.Equal(t, alice, black.String())
	require.True(t, found)
	require.Nil(t, err)
}

func TestGetPlayerAddressBlackIncorrect(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "notanaddress"
	black, found, err := storedGame.GetPlayerAddress("b")
	require.Nil(t, black)
	require.False(t, found)
	require.EqualError(t, err, "black address is invalid: notanaddress: decoding bech32 failed: invalid separator index -1")
}

func TestGetPlayerAddressRedCorrect(t *testing.T) {
	storedGame := GetStoredGame1()
	red, found, err := storedGame.GetPlayerAddress("r")
	require.Equal(t, bob, red.String())
	require.True(t, found)
	require.Nil(t, err)
}

func TestGetPlayerAddressRedIncorrect(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "notanaddress"
	red, found, err := storedGame.GetPlayerAddress("r")
	require.Nil(t, red)
	require.False(t, found)
	require.EqualError(t, err, "red address is invalid: notanaddress: decoding bech32 failed: invalid separator index -1")
}

func TestGetPlayerAddressWhiteNotFound(t *testing.T) {
	storedGame := GetStoredGame1()
	white, found, err := storedGame.GetPlayerAddress("w")
	require.Nil(t, white)
	require.False(t, found)
	require.Nil(t, err)
}

func TestGetPlayerAddressAnyNotFound(t *testing.T) {
	storedGame := GetStoredGame1()
	white, found, err := storedGame.GetPlayerAddress("*")
	require.Nil(t, white)
	require.False(t, found)
	require.Nil(t, err)
}

func TestGetWinnerBlackCorrect(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Winner = "b"
	winner, found, err := storedGame.GetWinnerAddress()
	require.Equal(t, alice, winner.String())
	require.True(t, found)
	require.Nil(t, err)
}

func TestGetWinnerRedCorrect(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Winner = "r"
	winner, found, err := storedGame.GetWinnerAddress()
	require.Equal(t, bob, winner.String())
	require.True(t, found)
	require.Nil(t, err)
}

func TestGetWinnerNotYetCorrect(t *testing.T) {
	storedGame := GetStoredGame1()
	winner, found, err := storedGame.GetWinnerAddress()
	require.Nil(t, winner)
	require.False(t, found)
	require.Nil(t, err)
}

func TestGameValidateOk(t *testing.T) {
	storedGame := GetStoredGame1()
	require.NoError(t, storedGame.Validate())
}
