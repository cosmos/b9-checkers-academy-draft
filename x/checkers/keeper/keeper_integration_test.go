package keeper_test

import (
	"testing"

	checkersapp "github.com/b9lab/checkers/app"
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	transferEventCount            = 3 // As emitted by the bank
	createEventCount              = 8
	playEventCountFirst           = 8 // Extra "sender" attribute emitted by the bank
	playEventCountNext            = 7
	rejectEventCount              = 4
	rejectEventCountWithTransfer  = 5 // Extra "sender" attribute emitted by the bank
	forfeitEventCount             = 4
	forfeitEventCountWithTransfer = 5 // Extra "sender" attribute emitted by the bank
	alice                         = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob                           = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol                         = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
	balAlice                      = 50000000
	balBob                        = 20000000
	balCarol                      = 10000000
	foreignToken                  = "foreignToken"
	balTokenAlice                 = 5
	balTokenBob                   = 2
	balTokenCarol                 = 1
)

var (
	checkersModuleAddress string
)

type IntegrationTestSuite struct {
	suite.Suite

	app         *checkersapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestCheckersKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := checkersapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	checkersModuleAddress = app.AccountKeeper.GetModuleAddress(types.ModuleName).String()

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.CheckersKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.msgServer = keeper.NewMsgServerImpl(app.CheckersKeeper)
	suite.ctx = ctx
	suite.queryClient = queryClient
}

func makeBalance(address string, balance int64, balanceToken int64) banktypes.Balance {
	return banktypes.Balance{
		Address: address,
		Coins: sdk.Coins{
			sdk.Coin{
				Denom:  sdk.DefaultBondDenom,
				Amount: sdk.NewInt(balance),
			},
			sdk.Coin{
				Denom:  foreignToken,
				Amount: sdk.NewInt(balanceToken),
			},
		},
	}
}

func getBankGenesis() *banktypes.GenesisState {
	coins := []banktypes.Balance{
		makeBalance(alice, balAlice, balTokenAlice),
		makeBalance(bob, balBob, balTokenBob),
		makeBalance(carol, balCarol, balTokenCarol),
	}
	supply := banktypes.NewSupply(coins[0].Coins.Add(coins[1].Coins...).Add(coins[2].Coins...))

	state := banktypes.NewGenesisState(
		banktypes.DefaultParams(),
		coins,
		supply.GetTotal(),
		[]banktypes.Metadata{})

	return state
}

func (suite *IntegrationTestSuite) setupSuiteWithBalances() {
	suite.app.BankKeeper.InitGenesis(suite.ctx, getBankGenesis())
}

func (suite *IntegrationTestSuite) RequireBankBalance(expected int, atAddress string) {
	suite.RequireBankBalanceIn(expected, atAddress, sdk.DefaultBondDenom)
}

func (suite *IntegrationTestSuite) RequireBankBalanceIn(expected int, atAddress string, denom string) {
	sdkAdd, err := sdk.AccAddressFromBech32(atAddress)
	suite.Require().Nil(err, "Address %s failed to parse")
	suite.Require().Equal(
		int64(expected),
		suite.app.BankKeeper.GetBalance(suite.ctx, sdkAdd, denom).Amount.Int64())
}
