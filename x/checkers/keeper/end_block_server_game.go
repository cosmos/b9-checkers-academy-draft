package keeper

import (
	"context"
	"fmt"
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
	if !found {
		panic("NextGame not found")
	}

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
			if storedGame.MoveCount == 0 {
				storedGame.Winner = rules.NO_PLAYER.Color
				// No point in keeping a game that was never played
				k.RemoveStoredGame(ctx, storedGameId)
			} else {
				storedGame.Winner, found = opponents[storedGame.Turn]
				if !found {
					panic(fmt.Sprintf(types.ErrCannotFindWinnerByColor.Error(), storedGame.Turn))
				}
				if storedGame.MoveCount <= 1 {
					k.MustRefundWager(ctx, &storedGame)
				} else {
					k.MustPayWinnings(ctx, &storedGame)
					k.MustRegisterPlayerForfeit(ctx, &storedGame)
				}
				k.SetStoredGame(ctx, storedGame)
			}
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
					sdk.NewAttribute(sdk.AttributeKeyAction, types.ForfeitGameEventKey),
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
