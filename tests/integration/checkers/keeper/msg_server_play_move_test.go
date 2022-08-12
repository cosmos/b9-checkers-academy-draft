package keeper_test

import (
	"github.com/b9lab/checkers/x/checkers/testutil"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) setupSuiteWithOneGameForPlayMove() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
}

func (suite *IntegrationTestSuite) TestPlayMoveSavedGame() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	keeper := suite.app.CheckersKeeper
	systemInfo, found := keeper.GetSystemInfo(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.SystemInfo{
		NextId:        2,
		FifoHeadIndex: "1",
		FifoTailIndex: "1",
	}, systemInfo)
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	suite.Require().EqualValues(types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "r",
		Black:       bob,
		Red:         carol,
		MoveCount:   uint64(1),
		BeforeIndex: "-1",
		AfterIndex:  "-1",
		Deadline:    types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:      "*",
		Wager:       45,
	}, game1)
}

func (suite *IntegrationTestSuite) TestPlayMovePlayerPaid() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(45, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestPlayMovePlayerPaidEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   0,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestPlayMoveCannotPayFails() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   balCarol + 1,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "2",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.Require().Nil(playMoveResponse)
	suite.Require().Equal("black cannot pay the wager: 10000000stake is smaller than 10000001stake: insufficient funds", err.Error())
}

func (suite *IntegrationTestSuite) TestPlayMoveEmitted() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 6)

	playEvent := events[3]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "move-played",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: bob},
			{Key: "game-index", Value: "1"},
			{Key: "captured-x", Value: "-1"},
			{Key: "captured-y", Value: "-1"},
			{Key: "winner", Value: "*"},
			{Key: "board", Value: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"},
		},
	}, playEvent)

	transferEvent := events[5]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: checkersModuleAddress},
		{Key: "sender", Value: bob},
		{Key: "amount", Value: "45stake"},
	}, transferEvent.Attributes)
}

func (suite *IntegrationTestSuite) TestPlayMoveEmittedEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   0,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 6)

	playEvent := events[3]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "move-played",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: carol},
			{Key: "game-index", Value: "1"},
			{Key: "captured-x", Value: "-1"},
			{Key: "captured-y", Value: "-1"},
			{Key: "winner", Value: "*"},
			{Key: "board", Value: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"},
		},
	}, playEvent)

	transferEvent := events[5]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: checkersModuleAddress},
		{Key: "sender", Value: carol},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes)
}

func (suite *IntegrationTestSuite) TestPlayMove2PlayerPaid() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(45, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol-45, carol)
	suite.RequireBankBalance(90, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestPlayMove2CannotPayFails() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   balCarol + 1,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "2",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "2",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	suite.Require().Nil(playMoveResponse)
	suite.Require().Equal("red cannot pay the wager: 10000000stake is smaller than 10000001stake: insufficient funds", err.Error())
}

func (suite *IntegrationTestSuite) TestPlayMove3DidNotPay() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol-45, carol)
	suite.RequireBankBalance(90, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     2,
		FromY:     3,
		ToX:       0,
		ToY:       5,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol-45, carol)
	suite.RequireBankBalance(90, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestPlayMoveToWinnerBankPaid() {
	suite.setupSuiteWithOneGameForPlayMove()
	testutil.PlayAllMoves(suite.T(), suite.msgServer, sdk.WrapSDKContext(suite.ctx), "1", testutil.Game1Moves)
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob+45, bob)
	suite.RequireBankBalance(balCarol-45, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}
