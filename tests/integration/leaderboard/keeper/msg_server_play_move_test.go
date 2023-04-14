package keeper_test

import (
	"github.com/b9lab/checkers/x/checkers/testutil"
	checkerstypes "github.com/b9lab/checkers/x/checkers/types"
	leaderboardtypes "github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) setupSuiteWithOneGameForPlayMove() {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &checkerstypes.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   0,
		Denom:   "stake",
	})
}

func (suite *IntegrationTestSuite) TestPlayMoveToWinnerAddedToLeaderboard() {
	suite.setupSuiteWithOneGameForPlayMove()
	suite.app.CheckersKeeper.SetPlayerInfo(suite.ctx, checkerstypes.PlayerInfo{
		Index: alice, WonCount: 10,
	})
	suite.app.CheckersKeeper.SetPlayerInfo(suite.ctx, checkerstypes.PlayerInfo{
		Index: bob, WonCount: 10,
	})
	suite.app.LeaderboardKeeper.SetLeaderboard(suite.ctx, leaderboardtypes.Leaderboard{
		Winners: []leaderboardtypes.Winner{
			{Address: alice, WonCount: 10, AddedAt: 1000},
			{Address: bob, WonCount: 10, AddedAt: 999},
		},
	})
	testutil.PlayAllMoves(suite.T(), suite.msgServer, sdk.WrapSDKContext(suite.ctx), "1", bob, carol, testutil.Game1Moves)
	suite.app.LeaderboardKeeper.CollectSortAndClipLeaderboard(suite.ctx)
	leaderboard := suite.app.LeaderboardKeeper.GetLeaderboard(suite.ctx)
	suite.Require().EqualValues(
		[]leaderboardtypes.Winner{
			{Address: bob, WonCount: 11, AddedAt: uint64(suite.ctx.BlockTime().Unix())},
			{Address: alice, WonCount: 10, AddedAt: 1000},
		},
		leaderboard.Winners)
}
