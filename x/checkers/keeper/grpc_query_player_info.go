package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/xavierlepretre/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PlayerInfoAll(c context.Context, req *types.QueryAllPlayerInfoRequest) (*types.QueryAllPlayerInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var playerInfos []*types.PlayerInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	playerInfoStore := prefix.NewStore(store, types.KeyPrefix(types.PlayerInfoKey))

	pageRes, err := query.Paginate(playerInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var playerInfo types.PlayerInfo
		if err := k.cdc.UnmarshalBinaryBare(value, &playerInfo); err != nil {
			return err
		}

		playerInfos = append(playerInfos, &playerInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPlayerInfoResponse{PlayerInfo: playerInfos, Pagination: pageRes}, nil
}

func (k Keeper) PlayerInfo(c context.Context, req *types.QueryGetPlayerInfoRequest) (*types.QueryGetPlayerInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPlayerInfo(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPlayerInfoResponse{PlayerInfo: &val}, nil
}
