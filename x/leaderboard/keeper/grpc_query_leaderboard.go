package keeper

import (
	"context"

	"github.com/b9lab/checkers/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Leaderboard(c context.Context, req *types.QueryGetLeaderboardRequest) (*types.QueryGetLeaderboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val := k.GetLeaderboard(ctx)

	return &types.QueryGetLeaderboardResponse{Leaderboard: val}, nil
}
