package keeper

import (
	"context"
	"strings"

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
		if strings.Compare(storedGameId, types.NoFifoIdKey) == 0 {
			break
		}
		storedGame, found = k.GetStoredGame(ctx, storedGameId)
		if !found {
			panic("Fifo head game not found " + nextGame.FifoHead)
		}
		deadline, err := storedGame.GetDeadlineAsTime()
		if err != nil {
			panic(err)
		}
		if deadline.Before(ctx.BlockTime()) {
			// Game is past deadline
			k.RemoveFromFifo(ctx, &storedGame, &nextGame)
			fullGame := storedGame.ToFullGame()

			if fullGame.MoveCount == 0 {
				// No point in keeping a game that was never played
				k.RemoveStoredGame(ctx, storedGameId)
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(types.ForfeitGameEventKey,
						sdk.NewAttribute(types.ForfeitGameEventIdValue, storedGameId),
						sdk.NewAttribute(types.ForfeitGameEventWinner, rules.NO_PLAYER.Color),
					),
				)
			} else {
				fullGame.Winner, found = opponents[storedGame.Turn]
				if !found {
					panic("Could not find opponent of " + storedGame.Turn)
				}
				if fullGame.MoveCount <= 1 {
					k.MustRefundWager(ctx, &fullGame)
				} else {
					k.MustPayWinnings(ctx, &fullGame)
				}
				storedGame = fullGame.ToStoredGame()
				k.SetStoredGame(ctx, storedGame)
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(sdk.EventTypeMessage,
						sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
						sdk.NewAttribute(sdk.AttributeKeyAction, types.ForfeitGameEventKey),
						sdk.NewAttribute(types.ForfeitGameEventIdValue, storedGameId),
						sdk.NewAttribute(types.ForfeitGameEventWinner, storedGame.Winner),
					),
				)
			}
			// Move along FIFO
			storedGameId = nextGame.FifoHead
		} else {
			// All other games come after anyway
			break
		}
	}

	k.SetNextGame(ctx, nextGame)
}
