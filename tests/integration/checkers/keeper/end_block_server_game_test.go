package keeper_test

import (
	"time"

	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestForfeitUnplayedNotPaid() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))

	keeper.SetStoredGame(suite.ctx, game1)
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)

	keeper.ForfeitExpiredGames(goCtx)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitPlayedOnceRefunded() {
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
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(45, checkersModuleAddress)

	keeper.ForfeitExpiredGames(goCtx)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitPlayedOnceRefundedEmitted() {
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
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)

	keeper.ForfeitExpiredGames(goCtx)

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 7)

	forfeitEvent := events[2]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "game-forfeited",
		Attributes: []sdk.Attribute{
			{Key: "game-index", Value: "1"},
			{Key: "winner", Value: "*"},
			{Key: "board", Value: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"},
		},
	}, forfeitEvent)

	transferEvent := events[6]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "45stake"},
	}, transferEvent.Attributes[3:])
}

func (suite *IntegrationTestSuite) TestForfeitPlayedOnceRefundedEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   0,
	})
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	keeper := suite.app.CheckersKeeper
	game1, _ := keeper.GetStoredGame(suite.ctx, "1")
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)

	keeper.ForfeitExpiredGames(goCtx)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitPlayedOnceRefundedEmittedEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
		Wager:   0,
	})
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	keeper := suite.app.CheckersKeeper
	game1, _ := keeper.GetStoredGame(suite.ctx, "1")
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)

	keeper.ForfeitExpiredGames(goCtx)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 7)

	forfeitEvent := events[2]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "game-forfeited",
		Attributes: []sdk.Attribute{
			{Key: "game-index", Value: "1"},
			{Key: "winner", Value: "*"},
			{Key: "board", Value: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"},
		},
	}, forfeitEvent)

	transferEvent := events[6]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes[3:])
}

func (suite *IntegrationTestSuite) TestForfeitPlayedTwicePaid() {
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
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol-45, carol)
	suite.RequireBankBalance(90, checkersModuleAddress)

	keeper.ForfeitExpiredGames(goCtx)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-45, bob)
	suite.RequireBankBalance(balCarol+45, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitPlayedTwicePaidEmitted() {
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
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 7)

	forfeitEvent := events[2]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "game-forfeited",
		Attributes: []sdk.Attribute{
			{Key: "game-index", Value: "1"},
			{Key: "winner", Value: "r"},
			{Key: "board", Value: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*"},
		},
	}, forfeitEvent)

	transferEvent := events[6]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "90stake"},
	}, transferEvent.Attributes[6:])
}

func (suite *IntegrationTestSuite) TestForfeitOlderPlayedTwicePaidEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
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
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	keeper := suite.app.CheckersKeeper
	game1, _ := keeper.GetStoredGame(suite.ctx, "1")
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)

	keeper.ForfeitExpiredGames(goCtx)

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitOlderPlayedTwicePaidEmittedvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
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
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})
	keeper := suite.app.CheckersKeeper
	game1, _ := keeper.GetStoredGame(suite.ctx, "1")
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 7)

	forfeitEvent := events[2]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "game-forfeited",
		Attributes: []sdk.Attribute{
			{Key: "game-index", Value: "1"},
			{Key: "winner", Value: "r"},
			{Key: "board", Value: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*"},
		},
	}, forfeitEvent)

	transferEvent := events[6]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes[6:])
}
