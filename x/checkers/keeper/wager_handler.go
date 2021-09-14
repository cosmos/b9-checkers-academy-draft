package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

// Returns an error if the player has not enough funds.
func (k *Keeper) CollectWager(ctx sdk.Context, fullGame *types.FullGame) error {
	// Make the player pay the wager at the beginning
	if fullGame.MoveCount == 0 {
		// Black plays first
		err := k.bank.SendCoinsFromAccountToModule(ctx, fullGame.Black, types.ModuleName, sdk.NewCoins(fullGame.Wager))
		if err != nil {
			return errors.New("Black cannot pay the wager")
		}
	} else if fullGame.MoveCount == 1 {
		// Red plays second
		err := k.bank.SendCoinsFromAccountToModule(ctx, fullGame.Red, types.ModuleName, sdk.NewCoins(fullGame.Wager))
		if err != nil {
			return errors.New("Red cannot pay the wager")
		}
	}
	return nil
}

// Game must have a valid winner.
func (k *Keeper) MustPayWinnings(ctx sdk.Context, fullGame *types.FullGame) {
	// Pay the winnings to the winner
	winnerAddress, found := fullGame.GetWinnerAddress()
	if !found {
		panic("Could not get winner address by color: " + fullGame.Winner)
	}
	winnings := fullGame.Wager
	if fullGame.MoveCount == 0 {
		panic("There is nothing to pay, should not have been called")
	} else if 1 < fullGame.MoveCount {
		winnings = winnings.Add(winnings)
	}
	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winnerAddress, sdk.NewCoins(winnings))
	if err != nil {
		panic("Cannot pay the winnings to winner")
	}
}

// Game must be in a state where it can be refunded.
func (k *Keeper) MustRefundWager(ctx sdk.Context, fullGame *types.FullGame) {
	// Refund wager to black player if red rejects after black played
	if fullGame.MoveCount == 1 {
		err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fullGame.Black, sdk.NewCoins(fullGame.Wager))
		if err != nil {
			panic("Cannot refund the wager to black player")
		}
	} else if fullGame.MoveCount == 0 {
		// Do nothing
	} else {
		panic("Game is not in a state for refund")
	}
}
