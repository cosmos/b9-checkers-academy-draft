package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rules "github.com/xavierlepretre/checkers/x/checkers/rules"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func (k Keeper) ForfeitExpiredGames(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	opponents := map[string]string{
		rules.BLACK_PLAYER.Color: rules.RED_PLAYER.Color,
		rules.RED_PLAYER.Color:   rules.BLACK_PLAYER.Color,
	}

	// Get FIFO information
	nextGame, found := k.GetNextGame(ctx)

	storedGameId := nextGame.FifoHead
	var storedGame types.StoredGame
	for {
		// Finished moving along
		if storedGameId == types.NoFifoIdKey {
			break
		}
		storedGame, found = k.GetStoredGame(ctx, storedGameId)
		if !found {
			panic("Fifo head game not found " + nextGame.FifoHead)
		}
		if storedGame.ToFullGame().Deadline.Before(ctx.BlockTime()) {
			// Game is past deadline
			storedGame.Winner, found = opponents[storedGame.Turn]
			if !found {
				panic("Could not find opponent of " + storedGame.Turn)
			}
			k.RemoveFromFifo(ctx, &storedGame, &nextGame)
			k.SetStoredGame(ctx, storedGame)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(types.ForfeitGameEventKey,
					sdk.NewAttribute(types.ForfeitGameEventIdValue, storedGameId),
					sdk.NewAttribute(types.ForfeitGameEventWinner, storedGame.Winner),
				),
			)
			// Move along FIFO
			storedGameId = nextGame.FifoHead
		} else {
			// All other games come after anyway
			break
		}
	}

	k.SetNextGame(ctx, nextGame)
}
