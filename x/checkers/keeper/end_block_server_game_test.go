package keeper_test

import (
	"time"

	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestForfeitUnplayed() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitEvent.Attributes[createEventCount:])

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitOlderUnplayed() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  3,
		FifoHead: "2",
		FifoTail: "2",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitEvent.Attributes[2*createEventCount:])
}

func (suite *IntegrationTestSuite) TestForfeit2OldestUnplayedIn1Call() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   13,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	game2, found := keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().True(found)
	game2.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game2)
	keeper.ForfeitExpiredGames(goCtx)

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)
	_, found = keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  4,
		FifoHead: "3",
		FifoTail: "3",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	forfeitAttributes := forfeitEvent.Attributes[3*createEventCount:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitAttributes[:4])
	forfeitAttributes = forfeitAttributes[4:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "2"},
		{Key: "Winner", Value: "*"},
	}, forfeitAttributes)
}

func (suite *IntegrationTestSuite) TestForfeitPlayedOnce() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)
	suite.RequireBankBalance(balCarol, carol) // Refunded

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)
	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitEvent.Attributes[createEventCount+playEventCountFirst:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "11stake"},
	}, transferEvent.Attributes[transferEventCount:])

	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestForfeitOlderPlayedOnce() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)
	suite.RequireBankBalance(balCarol, carol) // Refunded

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  3,
		FifoHead: "2",
		FifoTail: "2",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitEvent.Attributes[2*createEventCount+playEventCountFirst:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "11stake"},
	}, transferEvent.Attributes[transferEventCount:])
}

func (suite *IntegrationTestSuite) TestForfeitOlderPlayedOncePaidEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   0,
	})
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	keeper := suite.app.CheckersKeeper
	game1, _ := keeper.GetStoredGame(suite.ctx, "1")
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)
	suite.RequireBankBalance(balCarol, carol) // Refunded

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitEvent.Attributes[createEventCount+playEventCountFirst:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes[transferEventCount:])
}

func (suite *IntegrationTestSuite) TestForfeit2OldestPlayedOnceIn1Call() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: alice,
		IdValue: "2",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   13,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	game1.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game1)
	game2, found := keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().True(found)
	game2.Deadline = types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(suite.ctx, game2)
	keeper.ForfeitExpiredGames(goCtx)
	suite.RequireBankBalance(balAlice, alice) // Refunded
	suite.RequireBankBalance(balCarol, carol) // Refunded

	_, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().False(found)
	_, found = keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().False(found)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  4,
		FifoHead: "3",
		FifoTail: "3",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	forfeitAttributes := forfeitEvent.Attributes[3*createEventCount+2*playEventCountFirst:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "*"},
	}, forfeitAttributes[:5])
	forfeitAttributes = forfeitAttributes[5:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "2"},
		{Key: "Winner", Value: "*"},
	}, forfeitAttributes)

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	transferAttributes := transferEvent.Attributes[2*transferEventCount:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "11stake"},
	}, transferAttributes[:3])
	transferAttributes = transferAttributes[3:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: alice},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "12stake"},
	}, transferAttributes)
}

func (suite *IntegrationTestSuite) TestForfeitPlayedTwice() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.RequireBankBalance(balBob, bob)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)
	suite.RequireBankBalance(balBob+11, bob)     // Won wager
	suite.RequireBankBalance(balCarol-11, carol) // Lost wager

	game1, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(2),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  oldDeadline,
		Winner:    "r",
		Wager:     11,
	}, game1)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "-1",
		FifoTail: "-1",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "r"},
	}, forfeitEvent.Attributes[createEventCount+2*playEventCountFirst:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "22stake"},
	}, transferEvent.Attributes[2*transferEventCount:])
}

func (suite *IntegrationTestSuite) TestForfeitOlderPlayedTwice() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.RequireBankBalance(balBob, bob)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)
	suite.RequireBankBalance(balBob+11, bob)     // Won wager
	suite.RequireBankBalance(balCarol-11, carol) // Lost wager

	game1, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(2),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  oldDeadline,
		Winner:    "r",
		Wager:     11,
	}, game1)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  3,
		FifoHead: "2",
		FifoTail: "2",
	}, nextGame)
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "r"},
	}, forfeitEvent.Attributes[2*createEventCount+2*playEventCountFirst:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "22stake"},
	}, transferEvent.Attributes[2*transferEventCount:])
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
	suite.RequireBankBalance(balCarol, carol)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.RequireBankBalance(balBob, bob)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	keeper := suite.app.CheckersKeeper
	game1, _ := keeper.GetStoredGame(suite.ctx, "1")
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)
	keeper.ForfeitExpiredGames(goCtx)
	suite.RequireBankBalance(balBob, bob)     // No wagers to win
	suite.RequireBankBalance(balCarol, carol) // No wagers to lose

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "r"},
	}, forfeitEvent.Attributes[createEventCount+2*playEventCountFirst:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes[2*transferEventCount:])
}

func (suite *IntegrationTestSuite) TestForfeit2OldestPlayedTwiceIn1Call() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(11, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(22, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: alice,
		IdValue: "2",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "2",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.RequireBankBalance(46, checkersModuleAddress)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   13,
	})
	keeper := suite.app.CheckersKeeper
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)
	game2, found := keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().True(found)
	game2.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game2)
	keeper.ForfeitExpiredGames(goCtx)

	game1, found = keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(2),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  oldDeadline,
		Winner:    "r",
		Wager:     11,
	}, game1)

	game2, found = keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().True(found)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   bob,
		Index:     "2",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       carol,
		Black:     alice,
		MoveCount: uint64(2),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  oldDeadline,
		Winner:    "r",
		Wager:     12,
	}, game2)

	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  4,
		FifoHead: "3",
		FifoTail: "3",
	}, nextGame)

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	forfeitEvent := events[0]
	suite.Require().Equal(forfeitEvent.Type, "message")
	forfeitAttributes := forfeitEvent.Attributes[3*createEventCount+4*playEventCountFirst:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "1"},
		{Key: "Winner", Value: "r"},
	}, forfeitAttributes[:5])
	forfeitAttributes = forfeitAttributes[5:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "GameForfeited"},
		{Key: "IdValue", Value: "2"},
		{Key: "Winner", Value: "r"},
	}, forfeitAttributes)

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	transferAttributes := transferEvent.Attributes[4*transferEventCount:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: bob},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "22stake"},
	}, transferAttributes[:3])
	transferAttributes = transferAttributes[3:]
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: carol},
		{Key: "sender", Value: checkersModuleAddress},
		{Key: "amount", Value: "24stake"},
	}, transferAttributes)

	suite.RequireBankBalance(balAlice-12, alice) // Lost wager
	suite.RequireBankBalance(balBob+11, bob)     // Won wager
	suite.RequireBankBalance(balCarol+1, carol)  // Lost and won wagers
	suite.RequireBankBalance(0, checkersModuleAddress)
}
