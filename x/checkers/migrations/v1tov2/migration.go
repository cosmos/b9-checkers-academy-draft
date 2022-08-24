package v1tov2

import (
	"github.com/b9lab/checkers/x/checkers/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func PerformMigration(ctx sdk.Context, k keeper.Keeper, storedGameChunk uint64, playerInfoChunk uint64) error {
	ctx.Logger().Info("Start to compute checkers games to player info calculation...")
	err := MapStoredGamesReduceToPlayerInfo(ctx, k, storedGameChunk)
	if err != nil {
		return err
	}
	ctx.Logger().Info("Checkers games to player info computation done")
	ctx.Logger().Info("Start to compute checkers player info to leaderboard calculation...")
	err = MapPlayerInfosReduceToLeaderboard(ctx, k, playerInfoChunk)
	if err != nil {
		return err
	}
	ctx.Logger().Info("Checkers player info to leaderboard computation done")
	return nil
}
