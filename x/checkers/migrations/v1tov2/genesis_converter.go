package v1tov2

import (
	"time"

	"github.com/b9lab/checkers/x/checkers/types"
)

func (genesisV1 GenesisStateV1) Convert(now time.Time) (genesis *types.GenesisState, err error) {
	playerInfos := make(map[string]*types.PlayerInfo, 1000)
	err = PopulatePlayerInfosWith(&playerInfos, &genesisV1.StoredGameList)
	if err != nil {
		return nil, err
	}
	leaderboard := CreateLeaderboardForGenesis()
	err = PopulateLeaderboardWith(leaderboard, &playerInfos, now)
	if err != nil {
		return nil, err
	}
	return &types.GenesisState{
		Leaderboard:    leaderboard,
		PlayerInfoList: PlayerInfoMapToList(&playerInfos),
		StoredGameList: genesisV1.StoredGameList,
		NextGame:       genesisV1.NextGame,
	}, nil
}
