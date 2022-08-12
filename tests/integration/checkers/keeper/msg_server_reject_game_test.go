package keeper_test

import (
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) setupSuiteWithOneGameForRejectGame() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})
}

func (suite *IntegrationTestSuite) TestRejectGameByBlackNoMove() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	rejectGameResponse, err := suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func (suite *IntegrationTestSuite) TestRejectGameByBlackNoMoveNotPaid() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedNoMoveNotPaid() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator:   bob,
		GameIndex: "1",
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveRefunded() {
	suite.setupSuiteWithOneGameForRejectGame()
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
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveEmitted() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 7)

	rejectEvent := events[2]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "game-rejected",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: carol},
			{Key: "game-index", Value: "1"},
		},
	}, rejectEvent)

	transferEvent := events[6]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "45stake"},
	}, transferEvent.Attributes[3:])
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   0,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveEvenZeroEmitted() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   0,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator:   carol,
		GameIndex: "1",
	})

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 7)

	rejectEvent := events[2]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "game-rejected",
		Attributes: []sdk.Attribute{
			{Key: "creator", Value: carol},
			{Key: "game-index", Value: "1"},
		},
	}, rejectEvent)

	transferEvent := events[6]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes[3:])
}
