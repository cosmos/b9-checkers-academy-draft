package v1tov2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/rules"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func getOrNewPlayerInfo(infoSoFar *map[string]*types.PlayerInfo, playerIndex string) (playerInfo *types.PlayerInfo) {
	playerInfo, found := (*infoSoFar)[playerIndex]
	if !found {
		playerInfo = &types.PlayerInfo{
			Index:          playerIndex,
			WonCount:       0,
			LostCount:      0,
			ForfeitedCount: 0,
		}
	}
	return playerInfo
}

func PopulatePlayerInfosWith(infoSoFar *map[string]*types.PlayerInfo, games *[]*types.StoredGame) (err error) {
	var winnerAddress, loserAddress sdk.AccAddress
	var winnerIndex, loserIndex string
	var winnerInfo, loserInfo *types.PlayerInfo
	for _, game := range *games {
		winnerAddress, err = game.GetRedAddress()
		if err != nil {
			return err
		}
		loserAddress, err = game.GetBlackAddress()
		if err != nil {
			return err
		}
		if game.Winner == rules.RED_PLAYER.Color {
			// Already correct
		} else if game.Winner == rules.BLACK_PLAYER.Color {
			winnerAddress, loserAddress = loserAddress, winnerAddress
		} else {
			// Game is still unresolved.
			continue
		}
		winnerIndex = winnerAddress.String()
		loserIndex = loserAddress.String()
		winnerInfo = getOrNewPlayerInfo(infoSoFar, winnerIndex)
		loserInfo = getOrNewPlayerInfo(infoSoFar, loserIndex)
		winnerInfo.WonCount += 1
		loserInfo.LostCount += 1
		(*infoSoFar)[winnerIndex] = winnerInfo
		(*infoSoFar)[loserIndex] = loserInfo
	}
	return nil
}

func PlayerInfoMapToList(all *map[string]*types.PlayerInfo) []*types.PlayerInfo {
	asList := make([]*types.PlayerInfo, 0, len(*all))
	for _, playerInfo := range *all {
		asList = append(asList, playerInfo)
	}
	return asList
}
