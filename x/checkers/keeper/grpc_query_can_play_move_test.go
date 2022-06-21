package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/b9lab/checkers/x/checkers/types"
)

type canPlayBoard struct {
	desc     string
	board    string
	turn     string
	request  *types.QueryCanPlayMoveRequest
	response *types.QueryCanPlayMoveResponse
	err      error
}

var (
	canPlayOkResponse = &types.QueryCanPlayMoveResponse{
		Possible: true,
		Reason:   "ok",
	}
	firstTestCase = canPlayBoard{
		desc:  "First move by black",
		board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		turn:  "b",
		request: &types.QueryCanPlayMoveRequest{
			IdValue: "1",
			Player:  "b",
			FromX:   1,
			FromY:   2,
			ToX:     2,
			ToY:     3,
		},
		response: canPlayOkResponse,
	}
	canPlayTestRange = []canPlayBoard{
		firstTestCase,
		{
			desc:  "First move by red, wrong",
			board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
			turn:  "b",
			request: &types.QueryCanPlayMoveRequest{
				IdValue: "1",
				Player:  "r",
				FromX:   1,
				FromY:   2,
				ToX:     2,
				ToY:     3,
			},
			response: &types.QueryCanPlayMoveResponse{
				Possible: false,
				Reason:   "player tried to play out of turn: red",
			},
		},
		{
			desc:  "Black can win",
			board: "*b*b****|**b*b***|*****b**|********|********|**r*****|*B***b**|********",
			turn:  "b",
			request: &types.QueryCanPlayMoveRequest{
				IdValue: "1",
				Player:  "b",
				FromX:   1,
				FromY:   6,
				ToX:     3,
				ToY:     4,
			},
			response: canPlayOkResponse,
		},
		{
			desc:  "Black must capture, see next for right move",
			board: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
			turn:  "b",
			request: &types.QueryCanPlayMoveRequest{
				IdValue: "1",
				Player:  "b",
				FromX:   7,
				FromY:   2,
				ToX:     6,
				ToY:     3,
			},
			response: &types.QueryCanPlayMoveResponse{
				Possible: false,
				Reason:   "wrong move%!(EXTRA string=Invalid move: {7 2} to {6 3})",
			},
		},
		{
			desc:  "Black can capture, same board as previous",
			board: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
			turn:  "b",
			request: &types.QueryCanPlayMoveRequest{
				IdValue: "1",
				Player:  "b",
				FromX:   2,
				FromY:   3,
				ToX:     0,
				ToY:     5,
			},
			response: canPlayOkResponse,
		},
		{
			desc:  "Black king can capture backwards",
			board: "*b*b***b|**b*b***|***b***r|********|***r****|********|***r****|r*B*r*r*",
			turn:  "b",
			request: &types.QueryCanPlayMoveRequest{
				IdValue: "1",
				Player:  "b",
				FromX:   2,
				FromY:   7,
				ToX:     4,
				ToY:     5,
			},
			response: canPlayOkResponse,
		},
	}
)

func TestCanPlayAsExpected(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	for _, testCase := range canPlayTestRange {
		t.Run(testCase.desc, func(t *testing.T) {
			keeper.SetStoredGame(ctx, types.StoredGame{
				Creator:   alice,
				Index:     testCase.request.IdValue,
				Game:      testCase.board,
				Turn:      testCase.turn,
				Red:       bob,
				Black:     carol,
				MoveCount: 1,
				BeforeId:  "-1",
				AfterId:   "-1",
				Deadline:  "",
				Winner:    "*",
				Wager:     0,
			})
			response, err := keeper.CanPlayMove(goCtx, testCase.request)
			require.Nil(t, err)
			require.EqualValues(t, testCase.response, response)
		})
	}
}

func TestCanPlayWrongNoRequest(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	_, err := keeper.CanPlayMove(goCtx, nil)
	require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
}

func TestCanPlayWrongNoGame(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	keeper.SetStoredGame(ctx, types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      firstTestCase.board,
		Turn:      firstTestCase.turn,
		Red:       bob,
		Black:     carol,
		MoveCount: 1,
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  "",
		Winner:    "*",
		Wager:     0,
	})
	_, err := keeper.CanPlayMove(goCtx, &types.QueryCanPlayMoveRequest{
		IdValue: "2",
		Player:  "b",
		FromX:   2,
		FromY:   7,
		ToX:     4,
		ToY:     5,
	})
	require.NotNil(t, err)
	require.EqualError(t, err, "game by id not found: 2: game by id not found: %s")
}

func (suite *IntegrationTestSuite) TestCanPlayAfterCreate() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	response, err := suite.queryClient.CanPlayMove(goCtx, firstTestCase.request)
	suite.Require().Nil(err)
	suite.Require().EqualValues(firstTestCase.response, response)
}
