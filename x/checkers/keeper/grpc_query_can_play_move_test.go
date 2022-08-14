package keeper_test

import (
	"testing"

	keepertest "github.com/b9lab/checkers/testutil/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type canPlayGameCase struct {
	desc     string
	game     types.StoredGame
	request  *types.QueryCanPlayMoveRequest
	response *types.QueryCanPlayMoveResponse
	err      string
}

var (
	canPlayOkResponse = &types.QueryCanPlayMoveResponse{
		Possible: true,
		Reason:   "ok",
	}
	canPlayTestRange = []canPlayGameCase{
		{
			desc: "First move by black",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "b",
				FromX:     1,
				FromY:     2,
				ToX:       2,
				ToY:       3,
			},
			response: canPlayOkResponse,
			err:      "nil",
		},
		{
			desc: "Nil request, wrong",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request:  nil,
			response: nil,
			err:      "rpc error: code = InvalidArgument desc = invalid request",
		},
		{
			desc: "Unknown game, wrong",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "2",
				Player:    "b",
				FromX:     1,
				FromY:     2,
				ToX:       2,
				ToY:       3,
			},
			response: nil,
			err:      "2: game by id not found",
		},
		{
			desc: "Game finished, wrong",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b****|**b*b***|*****b**|********|***B****|********|*****b**|********",
				Turn:   "r",
				Winner: "b",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "u",
				FromX:     1,
				FromY:     2,
				ToX:       2,
				ToY:       3,
			},
			response: &types.QueryCanPlayMoveResponse{
				Possible: false,
				Reason:   "game is already finished",
			},
			err: "nil",
		},
		{
			desc: "Game not parseable, wrong",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "b",
				FromX:     1,
				FromY:     2,
				ToX:       2,
				ToY:       3,
			},
			response: nil,
			err:      "game cannot be parsed: invalid board string: *b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r",
		},
		{
			desc: "First move by unknown, wrong",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "u",
				FromX:     1,
				FromY:     2,
				ToX:       2,
				ToY:       3,
			},
			response: &types.QueryCanPlayMoveResponse{
				Possible: false,
				Reason:   "message creator is not a player: u",
			},
			err: "nil",
		},
		{
			desc: "First move by red, wrong",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "r",
				FromX:     1,
				FromY:     2,
				ToX:       2,
				ToY:       3,
			},
			response: &types.QueryCanPlayMoveResponse{
				Possible: false,
				Reason:   "player tried to play out of turn: red",
			},
			err: "nil",
		},
		{
			desc: "Black can win",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b****|**b*b***|*****b**|********|********|**r*****|*B***b**|********",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "b",
				FromX:     1,
				FromY:     6,
				ToX:       3,
				ToY:       4,
			},
			response: canPlayOkResponse,
			err:      "nil",
		},
		{
			desc: "Black must capture, see next for right move",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "b",
				FromX:     7,
				FromY:     2,
				ToX:       6,
				ToY:       3,
			},
			response: &types.QueryCanPlayMoveResponse{
				Possible: false,
				Reason:   "wrong move: Invalid move: {7 2} to {6 3}",
			},
			err: "nil",
		},
		{
			desc: "Black can capture, same board as previous",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|*r******|**r*r*r*|*r*r*r*r|r*r*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "b",
				FromX:     2,
				FromY:     3,
				ToX:       0,
				ToY:       5,
			},
			response: canPlayOkResponse,
			err:      "nil",
		},
		{
			desc: "Black king can capture backwards",
			game: types.StoredGame{
				Index:  "1",
				Board:  "*b*b***b|**b*b***|***b***r|********|***r****|********|***r****|r*B*r*r*",
				Turn:   "b",
				Winner: "*",
			},
			request: &types.QueryCanPlayMoveRequest{
				GameIndex: "1",
				Player:    "b",
				FromX:     2,
				FromY:     7,
				ToX:       4,
				ToY:       5,
			},
			response: canPlayOkResponse,
			err:      "nil",
		},
	}
)

func TestCanPlayCasesAsExpected(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	goCtx := sdk.WrapSDKContext(ctx)
	for _, testCase := range canPlayTestRange {
		t.Run(testCase.desc, func(t *testing.T) {
			keeper.SetStoredGame(ctx, testCase.game)
			response, err := keeper.CanPlayMove(goCtx, testCase.request)
			if testCase.response == nil {
				require.Nil(t, response)
			} else {
				require.EqualValues(t, testCase.response, response)
			}
			if testCase.err == "nil" {
				require.Nil(t, err)
			} else {
				require.EqualError(t, err, testCase.err)
			}
			keeper.RemoveStoredGame(ctx, testCase.game.Index)
		})
	}
}
