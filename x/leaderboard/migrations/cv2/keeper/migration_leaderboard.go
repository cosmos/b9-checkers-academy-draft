package keeper

import (
	"context"
	"time"

	checkerstypes "github.com/b9lab/checkers/x/checkers/types"
	"github.com/b9lab/checkers/x/leaderboard/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

type PlayerInfosChunk struct {
	PlayerInfos []checkerstypes.PlayerInfo
	Error       error
}

func LoadPlayerInfosToChannel(context context.Context,
	playerInfosk types.PlayerInfoKeeper,
	playerInfosChannel chan<- PlayerInfosChunk,
	chunk uint64) {
	defer func() { close(playerInfosChannel) }()
	response, err := playerInfosk.PlayerInfoAll(context, &checkerstypes.QueryAllPlayerInfoRequest{
		Pagination: &query.PageRequest{Limit: chunk},
	})
	if err != nil {
		playerInfosChannel <- PlayerInfosChunk{PlayerInfos: nil, Error: err}
		return
	}
	playerInfosChannel <- PlayerInfosChunk{PlayerInfos: response.PlayerInfo, Error: nil}

	for response.Pagination.NextKey != nil {
		response, err = playerInfosk.PlayerInfoAll(context, &checkerstypes.QueryAllPlayerInfoRequest{
			Pagination: &query.PageRequest{
				Key:   response.Pagination.NextKey,
				Limit: chunk,
			},
		})
		if err != nil {
			playerInfosChannel <- PlayerInfosChunk{PlayerInfos: nil, Error: err}
			return
		}
		playerInfosChannel <- PlayerInfosChunk{PlayerInfos: response.PlayerInfo, Error: nil}
	}
}

type LeaderboardResult struct {
	Leaderboard *types.Leaderboard
	Error       error
}

func HandlePlayerInfosChannel(playerInfosChannel <-chan PlayerInfosChunk,
	leaderboardChannel chan<- LeaderboardResult,
	leaderboardLength uint64,
	addedAt time.Time,
	chunk uint64) {
	defer func() { close(leaderboardChannel) }()
	winners := make([]types.Winner, 0, leaderboardLength+chunk)
	for receivedInfos := range playerInfosChannel {
		if receivedInfos.Error != nil {
			leaderboardChannel <- LeaderboardResult{Leaderboard: nil, Error: receivedInfos.Error}
			return
		}
		if receivedInfos.PlayerInfos != nil {
			candidates, err := types.MakeCandidatesFromPlayerInfos(receivedInfos.PlayerInfos)
			if err != nil {
				leaderboardChannel <- LeaderboardResult{Leaderboard: nil, Error: err}
				return
			}
			winners = types.AddCandidatesAtNow(winners, addedAt, candidates)
			if leaderboardLength < uint64(len(winners)) {
				winners = winners[:leaderboardLength]
			}
		}
	}
	leaderboardChannel <- LeaderboardResult{
		Leaderboard: &types.Leaderboard{Winners: winners},
		Error:       nil,
	}
}

func MapPlayerInfosReduceToLeaderboard(context context.Context,
	playerInfosk types.PlayerInfoKeeper,
	leaderboardLength uint64,
	addedAt time.Time,
	chunk uint64) (*types.Leaderboard, error) {
	playerInfosChannel := make(chan PlayerInfosChunk)
	leaderboardChannel := make(chan LeaderboardResult)

	go HandlePlayerInfosChannel(playerInfosChannel, leaderboardChannel, leaderboardLength, addedAt, chunk)
	go LoadPlayerInfosToChannel(context, playerInfosk, playerInfosChannel, chunk)

	result := <-leaderboardChannel

	return result.Leaderboard, result.Error
}
