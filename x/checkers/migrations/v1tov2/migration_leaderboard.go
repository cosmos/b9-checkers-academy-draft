package v1tov2

import (
	"time"

	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func addParsedCandidatesAndSort(parsedWinners []types.WinningPlayerParsed, candidates []types.WinningPlayerParsed) []types.WinningPlayerParsed {
	updated := append(parsedWinners, candidates...)
	types.SortWinners(updated)
	if types.LeaderboardWinnerLength < uint64(len(updated)) {
		updated = updated[:types.LeaderboardWinnerLength]
	}
	return updated
}

func AddCandidatesAndSortAtNow(parsedWinners []types.WinningPlayerParsed, now time.Time, playerInfos []types.PlayerInfo) []types.WinningPlayerParsed {
	parsedPlayers := make([]types.WinningPlayerParsed, 0, len(playerInfos))
	for _, playerInfo := range playerInfos {
		if playerInfo.WonCount > 0 {
			parsedPlayers = append(parsedPlayers, types.WinningPlayerParsed{
				PlayerAddress: playerInfo.Index,
				WonCount:      playerInfo.WonCount,
				DateAdded:     now,
			})
		}
	}
	return addParsedCandidatesAndSort(parsedWinners, parsedPlayers)
}

func AddCandidatesAndSort(parsedWinners []types.WinningPlayerParsed, ctx sdk.Context, playerInfos []types.PlayerInfo) []types.WinningPlayerParsed {
	return AddCandidatesAndSortAtNow(parsedWinners, types.GetDateAdded(ctx), playerInfos)
}

func handlePlayerInfosChannel(ctx sdk.Context, k keeper.Keeper,
	playerInfosChannel <-chan []types.PlayerInfo,
	done chan<- bool,
	chunk uint64) {
	winners := make([]types.WinningPlayerParsed, 0, types.LeaderboardWinnerLength+chunk)
	for receivedInfo := range playerInfosChannel {
		if receivedInfo != nil {
			winners = AddCandidatesAndSort(winners, ctx, receivedInfo)
		}
	}
	k.SetLeaderboard(ctx, types.CreateLeaderboardFromParsedWinners(winners))
	done <- true
}

func MapPlayerInfosReduceToLeaderboard(ctx sdk.Context, k keeper.Keeper, chunk uint64) error {
	context := sdk.WrapSDKContext(ctx)
	response, err := k.PlayerInfoAll(context, &types.QueryAllPlayerInfoRequest{
		Pagination: &query.PageRequest{
			Limit: chunk,
		},
	})
	if err != nil {
		return err
	}
	playerInfosChannel := make(chan []types.PlayerInfo)
	done := make(chan bool)

	go handlePlayerInfosChannel(ctx, k, playerInfosChannel, done, chunk)
	playerInfosChannel <- response.PlayerInfo

	for response.Pagination.NextKey != nil {
		response, err = k.PlayerInfoAll(context, &types.QueryAllPlayerInfoRequest{
			Pagination: &query.PageRequest{
				Key:   response.Pagination.NextKey,
				Limit: chunk,
			},
		})
		if err != nil {
			return err
		}
		playerInfosChannel <- response.PlayerInfo
	}
	close(playerInfosChannel)

	<-done
	return nil
}
