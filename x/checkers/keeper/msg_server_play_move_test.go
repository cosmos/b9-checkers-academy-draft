package keeper_test

import (
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) setupSuiteWithOneGameForPlayMove() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   11,
		Token:   sdk.DefaultBondDenom,
	})
}

func (suite *IntegrationTestSuite) TestPlayMove() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *IntegrationTestSuite) TestPlayMoveSameBlackRed() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   bob,
		Wager:   11,
		Token:   sdk.DefaultBondDenom,
	})
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *IntegrationTestSuite) TestPlayMovePlayerPaid() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
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
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol-11, carol)
	suite.RequireBankBalance(11, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestPlayMovePlayerPaidForeignToken() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   1,
		Token:   foreignToken,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.RequireBankBalanceIn(balTokenAlice, alice, foreignToken)
	suite.RequireBankBalanceIn(balTokenBob, bob, foreignToken)
	suite.RequireBankBalanceIn(balTokenCarol, carol, foreignToken)
	suite.RequireBankBalanceIn(0, checkersModuleAddress, foreignToken)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "2",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.RequireBankBalanceIn(balTokenAlice, alice, foreignToken)
	suite.RequireBankBalanceIn(balTokenBob, bob, foreignToken)
	suite.RequireBankBalanceIn(balTokenCarol-1, carol, foreignToken)
	suite.RequireBankBalanceIn(1, checkersModuleAddress, foreignToken)
}

func (suite *IntegrationTestSuite) TestPlayMoveConsumedGas() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	gasBefore := suite.ctx.GasMeter().GasConsumed()
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	gasAfter := suite.ctx.GasMeter().GasConsumed()
	suite.Require().Equal(uint64(33_461+10), gasAfter-gasBefore)
}

func (suite *IntegrationTestSuite) TestPlayMovePlayerPaidEvenZero() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   0,
		Token:   sdk.DefaultBondDenom,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
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
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)

	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	playEvent := events[0]
	suite.Require().Equal(playEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: carol},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "MovePlayed"},
		{Key: "Creator", Value: carol},
		{Key: "IdValue", Value: "1"},
		{Key: "CapturedX", Value: "-1"},
		{Key: "CapturedY", Value: "-1"},
		{Key: "Winner", Value: "*"},
	}, playEvent.Attributes[createEventCount:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: checkersModuleAddress},
		{Key: "sender", Value: carol},
		{Key: "amount", Value: ""},
	}, transferEvent.Attributes)
}

func (suite *IntegrationTestSuite) TestPlayMoveGasConsumedNoWager() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   0,
		Token:   sdk.DefaultBondDenom,
	})
	gasBefore := suite.ctx.GasMeter().GasConsumed()
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	gasAfter := suite.ctx.GasMeter().GasConsumed()
	suite.Require().Equal(uint64(26_534+10), gasAfter-gasBefore)
}

func (suite *IntegrationTestSuite) TestPlayMoveCannotPayFails() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   balCarol + 1,
		Token:   sdk.DefaultBondDenom,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "2",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	suite.Require().Nil(playMoveResponse)
	suite.Require().Equal("black cannot pay the wager: 10000000stake is smaller than 10000001stake: insufficient funds", err.Error())
}

func (suite *IntegrationTestSuite) TestPlayMoveSavedGame() {
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
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextGame)
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "r",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(1),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     11,
		Token:     "stake",
	}, game1)
}

func (suite *IntegrationTestSuite) TestPlayMoveWrongOutOfTurn() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.Require().Nil(playMoveResponse)
	suite.Require().Equal("player tried to play out of turn: %s", err.Error())
}

func (suite *IntegrationTestSuite) TestPlayMoveWrongPieceAtDestination() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   1,
		FromY:   0,
		ToX:     0,
		ToY:     1,
	})
	suite.Require().Nil(playMoveResponse)
	suite.Require().Equal("Already piece at destination position: {0 1}: wrong move", err.Error())
}

func (suite *IntegrationTestSuite) TestPlayMoveEmitted() {
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
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 2)

	playEvent := events[0]
	suite.Require().Equal(playEvent.Type, "message")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "sender", Value: carol},
		{Key: "module", Value: "checkers"},
		{Key: "action", Value: "MovePlayed"},
		{Key: "Creator", Value: carol},
		{Key: "IdValue", Value: "1"},
		{Key: "CapturedX", Value: "-1"},
		{Key: "CapturedY", Value: "-1"},
		{Key: "Winner", Value: "*"},
	}, playEvent.Attributes[createEventCount:])

	transferEvent := events[1]
	suite.Require().Equal(transferEvent.Type, "transfer")
	suite.Require().EqualValues([]sdk.Attribute{
		{Key: "recipient", Value: checkersModuleAddress},
		{Key: "sender", Value: carol},
		{Key: "amount", Value: "11stake"},
	}, transferEvent.Attributes)
}

func (suite *IntegrationTestSuite) TestPlayMove2() {
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
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *IntegrationTestSuite) TestPlayMove2PlayerPaid() {
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
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol-11, carol)
	suite.RequireBankBalance(11, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-11, bob)
	suite.RequireBankBalance(balCarol-11, carol)
	suite.RequireBankBalance(22, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestPlayMove2CannotPayFails() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     carol,
		Black:   bob,
		Wager:   balCarol + 1,
		Token:   sdk.DefaultBondDenom,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "2",
		FromX:   1,
		FromY:   2,
		ToX:     2,
		ToY:     3,
	})
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "2",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.Require().Nil(playMoveResponse)
	suite.Require().Equal("red cannot pay the wager: 10000000stake is smaller than 10000001stake: insufficient funds", err.Error())
}

func (suite *IntegrationTestSuite) TestPlayMove2SavedGame() {
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
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextGame)
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
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
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     11,
		Token:     "stake",
	}, game1)
}

func (suite *IntegrationTestSuite) TestPlayMove3() {
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
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	playMoveResponse, err := suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   2,
		FromY:   3,
		ToX:     0,
		ToY:     5,
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgPlayMoveResponse{
		IdValue:   "1",
		CapturedX: 1,
		CapturedY: 4,
		Winner:    "*",
	}, *playMoveResponse)
}

func (suite *IntegrationTestSuite) TestPlayMove3DidNotPay() {
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
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-11, bob)
	suite.RequireBankBalance(balCarol-11, carol)
	suite.RequireBankBalance(22, checkersModuleAddress)
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   2,
		FromY:   3,
		ToX:     0,
		ToY:     5,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob-11, bob)
	suite.RequireBankBalance(balCarol-11, carol)
	suite.RequireBankBalance(22, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestPlayMove3SavedGame() {
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
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: bob,
		IdValue: "1",
		FromX:   0,
		FromY:   5,
		ToX:     1,
		ToY:     4,
	})
	suite.msgServer.PlayMove(goCtx, &types.MsgPlayMove{
		Creator: carol,
		IdValue: "1",
		FromX:   2,
		FromY:   3,
		ToX:     0,
		ToY:     5,
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
	}, nextGame)
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|***b*b*b|********|********|b*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "r",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(3),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     11,
		Token:     "stake",
	}, game1)
}
