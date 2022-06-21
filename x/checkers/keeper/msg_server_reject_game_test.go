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
		Red:     bob,
		Black:   carol,
		Wager:   11,
	})
}

func (suite *IntegrationTestSuite) TestRejectGameWrongByCreator() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	rejectGameResponse, err := suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: alice,
		IdValue: "1",
	})
	suite.Require().Nil(rejectGameResponse)
	suite.Require().Equal("message creator is not a player: %s", err.Error())
}

func (suite *IntegrationTestSuite) TestRejectGameByBlackNoMove() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	rejectGameResponse, err := suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func (suite *IntegrationTestSuite) TestRejectGameByBlackNotPaid() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestRejectGameByBlackNoMoveRemovedGame() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)
}

func (suite *IntegrationTestSuite) TestRejectGameByBlackNoMoveEmitted() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)
	event := events[0]
	suite.Require().Equal(event.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameRejected"},
		{Key: "Creator", Value: carol},
		{Key: "IdValue", Value: "1"},
	}, event.Attributes[createEventCount:])
}

func (suite *IntegrationTestSuite) TestRejectGameByRedNoMove() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	rejectGameResponse, err := suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedNoMoveNotPaid() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedNoMoveRemovedGame() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedNoMoveEmitted() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)
	event := events[0]
	suite.Require().Equal(event.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameRejected"},
		{Key: "Creator", Value: bob},
		{Key: "IdValue", Value: "1"},
	}, event.Attributes[createEventCount:])
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMove() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	rejectGameResponse, err := suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgRejectGameResponse{}, *rejectGameResponse)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveRefunded() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol-11, carol)
	suite.RequireBankBalance(11, checkersModuleAddress)
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveRemovedGame() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveEmitted() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	rejectEvent := events[0]
	suite.Require().Equal(rejectEvent.Type, "message")
	rejectAttributesDiscardCount := createEventCount + playEventCountFirst
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameRejected"},
		{Key: "Creator", Value: bob},
		{Key: "IdValue", Value: "1"},
	}, rejectEvent.Attributes[rejectAttributesDiscardCount:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "11stake"},
	}, transferEvent.Attributes[transferEventCount:])
}

func (suite *IntegrationTestSuite) TestRejectGameByRedOneMoveEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   0,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	rejectEvent := events[0]
	suite.Require().Equal(rejectEvent.Type, "message")
	rejectAttributesDiscardCount := createEventCount + playEventCountFirst
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameRejected"},
		{Key: "Creator", Value: bob},
		{Key: "IdValue", Value: "1"},
	}, rejectEvent.Attributes[rejectAttributesDiscardCount:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes[transferEventCount:])
}

func (suite *IntegrationTestSuite) TestRejectGameByBlackWrongOneMove() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	rejectGameResponse, err := suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	suite.Require().Nil(rejectGameResponse)
	suite.Require().Equal("black player has already played", err.Error())
}

func (suite *IntegrationTestSuite) TestRejectGameByRedWrong2Moves() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	rejectGameResponse, err := suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: bob,
		IdValue: "1",
	})
	suite.Require().Nil(rejectGameResponse)
	suite.Require().Equal("red player has already played", err.Error())
}
