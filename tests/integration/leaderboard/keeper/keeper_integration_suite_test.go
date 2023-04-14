package keeper_test

import (
	"testing"
	"time"

	checkersapp "github.com/b9lab/checkers/app"
	checkerskeeper "github.com/b9lab/checkers/x/checkers/keeper"
	checkerstypes "github.com/b9lab/checkers/x/checkers/types"
	"github.com/b9lab/checkers/x/leaderboard/testutil"
	"github.com/b9lab/checkers/x/leaderboard/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

type IntegrationTestSuite struct {
	suite.Suite

	app         *checkersapp.App
	msgServer   checkerstypes.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestLeaderboardKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := checkersapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.LeaderboardKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.msgServer = checkerskeeper.NewMsgServerImpl(app.CheckersKeeper)
	suite.ctx = ctx
	suite.queryClient = queryClient
}
