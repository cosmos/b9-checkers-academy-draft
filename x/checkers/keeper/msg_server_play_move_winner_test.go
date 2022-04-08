package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

type GameMoveTest struct {
	player string
	fromX  uint64
	fromY  uint64
	toX    uint64
	toY    uint64
}

var (
	game1moves = []GameMoveTest{
		{"b", 1, 2, 2, 3}, // "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"
		{"r", 0, 5, 1, 4}, // "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*"
		{"b", 2, 3, 0, 5}, // "*b*b*b*b|b*b*b*b*|***b*b*b|********|********|b*r*r*r*|*r*r*r*r|r*r*r*r*"
		{"r", 4, 5, 3, 4}, // "*b*b*b*b|b*b*b*b*|***b*b*b|********|***r****|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"b", 3, 2, 2, 3}, // "*b*b*b*b|b*b*b*b*|*****b*b|**b*****|***r****|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"r", 3, 4, 1, 2}, // "*b*b*b*b|b*b*b*b*|*r***b*b|********|********|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"b", 0, 1, 2, 3}, // "*b*b*b*b|**b*b*b*|*****b*b|**b*****|********|b*r***r*|*r*r*r*r|r*r*r*r*"
		{"r", 2, 5, 3, 4}, // "*b*b*b*b|**b*b*b*|*****b*b|**b*****|***r****|b*****r*|*r*r*r*r|r*r*r*r*"
		{"b", 2, 3, 4, 5}, // "*b*b*b*b|**b*b*b*|*****b*b|********|********|b***b*r*|*r*r*r*r|r*r*r*r*"
		{"r", 5, 6, 3, 4}, // "*b*b*b*b|**b*b*b*|*****b*b|********|***r****|b*****r*|*r*r***r|r*r*r*r*"
		{"b", 5, 2, 4, 3}, // "*b*b*b*b|**b*b*b*|*******b|****b***|***r****|b*****r*|*r*r***r|r*r*r*r*"
		{"r", 3, 4, 5, 2}, // "*b*b*b*b|**b*b*b*|*****r*b|********|********|b*****r*|*r*r***r|r*r*r*r*"
		{"b", 6, 1, 4, 3}, // "*b*b*b*b|**b*b***|*******b|****b***|********|b*****r*|*r*r***r|r*r*r*r*"
		{"r", 6, 5, 5, 4}, // "*b*b*b*b|**b*b***|*******b|****b***|*****r**|b*******|*r*r***r|r*r*r*r*"
		{"b", 4, 3, 6, 5}, // "*b*b*b*b|**b*b***|*******b|********|********|b*****b*|*r*r***r|r*r*r*r*"
		{"r", 7, 6, 5, 4}, // "*b*b*b*b|**b*b***|*******b|********|*****r**|b*******|*r*r****|r*r*r*r*"
		{"b", 7, 2, 6, 3}, // "*b*b*b*b|**b*b***|********|******b*|*****r**|b*******|*r*r****|r*r*r*r*"
		{"r", 5, 4, 7, 2}, // "*b*b*b*b|**b*b***|*******r|********|********|b*******|*r*r****|r*r*r*r*"
		{"b", 4, 1, 3, 2}, // "*b*b*b*b|**b*****|***b***r|********|********|b*******|*r*r****|r*r*r*r*"
		{"r", 3, 6, 4, 5}, // "*b*b*b*b|**b*****|***b***r|********|********|b***r***|*r******|r*r*r*r*"
		{"b", 5, 0, 4, 1}, // "*b*b***b|**b*b***|***b***r|********|********|b***r***|*r******|r*r*r*r*"
		{"r", 2, 7, 3, 6}, // "*b*b***b|**b*b***|***b***r|********|********|b***r***|*r*r****|r***r*r*"
		{"b", 0, 5, 2, 7}, // "*b*b***b|**b*b***|***b***r|********|********|****r***|***r****|r*B*r*r*"
		{"r", 4, 5, 3, 4}, // "*b*b***b|**b*b***|***b***r|********|***r****|********|***r****|r*B*r*r*"
		{"b", 2, 7, 4, 5}, // "*b*b***b|**b*b***|***b***r|********|***r****|****B***|********|r***r*r*"
		// Captures again
		{"b", 4, 5, 2, 3}, // "*b*b***b|**b*b***|***b***r|**B*****|********|********|********|r***r*r*"
		{"r", 6, 7, 5, 6}, // "*b*b***b|**b*b***|***b***r|**B*****|********|********|*****r**|r***r***"
		{"b", 2, 3, 3, 4}, // "*b*b***b|**b*b***|***b***r|********|***B****|********|*****r**|r***r***"
		{"r", 0, 7, 1, 6}, // "*b*b***b|**b*b***|***b***r|********|***B****|********|*r***r**|****r***"
		{"b", 3, 2, 4, 3}, // "*b*b***b|**b*b***|*******r|****b***|***B****|********|*r***r**|****r***"
		{"r", 7, 2, 6, 1}, // "*b*b***b|**b*b*r*|********|****b***|***B****|********|*r***r**|****r***"
		{"b", 7, 0, 5, 2}, // "*b*b****|**b*b***|*****b**|****b***|***B****|********|*r***r**|****r***"
		{"r", 1, 6, 2, 5}, // "*b*b****|**b*b***|*****b**|****b***|***B****|**r*****|*****r**|****r***"
		{"b", 3, 4, 1, 6}, // "*b*b****|**b*b***|*****b**|****b***|********|********|*B***r**|****r***"
		{"r", 4, 7, 3, 6}, // "*b*b****|**b*b***|*****b**|****b***|********|********|*B*r*r**|********"
		{"b", 4, 3, 3, 4}, // "*b*b****|**b*b***|*****b**|********|***b****|********|*B*r*r**|********"
		{"r", 5, 6, 4, 5}, // "*b*b****|**b*b***|*****b**|********|***b****|****r***|*B*r****|********"
		{"b", 3, 4, 5, 6}, // "*b*b****|**b*b***|*****b**|********|********|********|*B*r*b**|********"
		{"r", 3, 6, 2, 5}, // "*b*b****|**b*b***|*****b**|********|********|**r*****|*B***b**|********"
		{"b", 1, 6, 3, 4}, // "*b*b****|**b*b***|*****b**|********|***B****|********|*****b**|********"
	}
)

func getPlayer(color string) string {
	if color == "b" {
		return carol
	}
	return bob
}

func (suite *IntegrationTestSuite) TestPlayMoveUpToWinner() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)

	for _, move := range game1moves {
		_, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
			Creator: getPlayer(move.player),
			IdValue: "1",
			FromX:   move.fromX,
			FromY:   move.fromY,
			ToX:     move.toX,
			ToY:     move.toY,
		})
		suite.Require().Nil(err)
	}

	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)

	game1, found1 := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found1)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b****|**b*b***|*****b**|********|***B****|********|*****b**|********",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(len(game1moves)),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "b",
		Wager:     11,
	}, game1)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	winEvent := events[0]
	suite.Require().Equal(winEvent.Type, "message")
	winWttributesDiscardCount := createEventCount + 2*playEventCountFirst + 37*playEventCountNext
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "MovePlayed"},
		{Key: "Creator", Value: carol},
		{Key: "IdValue", Value: "1"},
		{Key: "CapturedX", Value: "2"},
		{Key: "CapturedY", Value: "5"},
		{Key: "Winner", Value: "b"},
	}, winEvent.Attributes[winWttributesDiscardCount:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	transferAttributesDiscardCount := 2 * transferEventCount
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "22stake"},
	}, transferEvent.Attributes[transferAttributesDiscardCount:])

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-11, bob)
	suite.RequireBankBalance(balCarol+11, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}
