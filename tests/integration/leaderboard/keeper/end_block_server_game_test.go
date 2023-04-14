package keeper_test

import (
	"time"

	"github.com/b9lab/checkers/x/checkers/types"
	leaderboardtypes "github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestForfeitPlayedTwiceCalledHooks() {
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
	keeper.SetPlayerInfo(suite.ctx, types.PlayerInfo{
		Index: bob, WonCount: 10,
	})
	keeper.SetPlayerInfo(suite.ctx, types.PlayerInfo{
		Index: carol, WonCount: 10,
	})
	suite.app.LeaderboardKeeper.SetLeaderboard(suite.ctx, leaderboardtypes.Leaderboard{
		Winners: []leaderboardtypes.Winner{
			{Address: bob, WonCount: 10, AddedAt: 1000},
			{Address: carol, WonCount: 10, AddedAt: 999},
		},
	})
	game1, found := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found)
	oldDeadline := types.FormatDeadline(suite.ctx.BlockTime().Add(time.Duration(-1)))
	game1.Deadline = oldDeadline
	keeper.SetStoredGame(suite.ctx, game1)

	keeper.ForfeitExpiredGames(goCtx)
	suite.app.LeaderboardKeeper.CollectSortAndClipLeaderboard(suite.ctx)
	leaderboard := suite.app.LeaderboardKeeper.GetLeaderboard(suite.ctx)
	suite.Require().EqualValues(
		[]leaderboardtypes.Winner{
			{Address: carol, WonCount: 11, AddedAt: uint64(suite.ctx.BlockTime().Unix())},
			{Address: bob, WonCount: 10, AddedAt: 1000},
		},
		leaderboard.Winners)
}
