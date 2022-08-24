package v1tov2

import (
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/rules"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func getOrNewPlayerInfoInMap(infoSoFar *map[string]*types.PlayerInfo, playerIndex string) (playerInfo *types.PlayerInfo) {
	playerInfo, found := (*infoSoFar)[playerIndex]
	if !found {
		playerInfo = &types.PlayerInfo{
			Index:          playerIndex,
			WonCount:       0,
			LostCount:      0,
			ForfeitedCount: 0,
		}
		(*infoSoFar)[playerIndex] = playerInfo
	}
	return playerInfo
}

func handleStoredGameChannel(ctx sdk.Context,
	k keeper.Keeper,
	gamesChannel <-chan []types.StoredGame,
	playerInfoChannel chan<- *types.PlayerInfo) {
	for games := range gamesChannel {
		playerInfos := make(map[string]*types.PlayerInfo, len(games))
		for _, game := range games {
			var winner string
			var loser string
			if game.Winner == rules.PieceStrings[rules.BLACK_PLAYER] {
				winner = game.Black
				loser = game.Red
			} else if game.Winner == rules.PieceStrings[rules.RED_PLAYER] {
				winner = game.Red
				loser = game.Black
			} else {
				continue
			}
			getOrNewPlayerInfoInMap(&playerInfos, winner).WonCount++
			getOrNewPlayerInfoInMap(&playerInfos, loser).LostCount++
		}
		for _, playerInfo := range playerInfos {
			playerInfoChannel <- playerInfo
		}
	}
	close(playerInfoChannel)
}

func handlePlayerInfoChannel(ctx sdk.Context, k keeper.Keeper,
	playerInfoChannel <-chan *types.PlayerInfo,
	done chan<- bool) {
	for receivedInfo := range playerInfoChannel {
		if receivedInfo != nil {
			existingInfo, found := k.GetPlayerInfo(ctx, receivedInfo.Index)
			if found {
				existingInfo.WonCount += receivedInfo.WonCount
				existingInfo.LostCount += receivedInfo.LostCount
				existingInfo.ForfeitedCount += receivedInfo.ForfeitedCount
			} else {
				existingInfo = *receivedInfo
			}
			k.SetPlayerInfo(ctx, existingInfo)
		}
	}
	done <- true
}

func MapStoredGamesReduceToPlayerInfo(ctx sdk.Context, k keeper.Keeper, chunk uint64) error {
	context := sdk.WrapSDKContext(ctx)
	response, err := k.StoredGameAll(context, &types.QueryAllStoredGameRequest{
		Pagination: &query.PageRequest{
			Limit: chunk,
		},
	})
	if err != nil {
		return err
	}
	gamesChannel := make(chan []types.StoredGame)
	playerInfoChannel := make(chan *types.PlayerInfo)
	done := make(chan bool)

	go handleStoredGameChannel(ctx, k, gamesChannel, playerInfoChannel)
	go handlePlayerInfoChannel(ctx, k, playerInfoChannel, done)
	gamesChannel <- response.StoredGame

	for response.Pagination.NextKey != nil {
		response, err = k.StoredGameAll(context, &types.QueryAllStoredGameRequest{
			Pagination: &query.PageRequest{
				Key:   response.Pagination.NextKey,
				Limit: chunk,
			},
		})
		if err != nil {
			return err
		}
		gamesChannel <- response.StoredGame
	}
	close(gamesChannel)

	<-done
	return nil
}
