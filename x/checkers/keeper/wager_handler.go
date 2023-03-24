package keeper

import (
	"fmt"

	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k *Keeper) CollectWager(ctx sdk.Context, storedGame *types.StoredGame) error {
	if storedGame.MoveCount == 0 {
		// Black plays first
		black, err := storedGame.GetBlackAddress()
		if err != nil {
			panic(err.Error())
		}
		err = k.bank.SendCoinsFromAccountToModule(ctx, black, types.ModuleName, sdk.NewCoins(storedGame.GetWagerCoin()))
		if err != nil {
			return sdkerrors.Wrapf(err, types.ErrBlackCannotPay.Error())
		}
	} else if storedGame.MoveCount == 1 {
		// Red plays second
		red, err := storedGame.GetRedAddress()
		if err != nil {
			panic(err.Error())
		}
		err = k.bank.SendCoinsFromAccountToModule(ctx, red, types.ModuleName, sdk.NewCoins(storedGame.GetWagerCoin()))
		if err != nil {
			return sdkerrors.Wrapf(err, types.ErrRedCannotPay.Error())
		}
	}
	return nil
}

func (k *Keeper) MustPayWinnings(ctx sdk.Context, storedGame *types.StoredGame) {
	winnerAddress, found, err := storedGame.GetWinnerAddress()
	if err != nil {
		panic(err.Error())
	}
	if !found {
		panic(fmt.Sprintf(types.ErrCannotFindWinnerByColor.Error(), storedGame.Winner))
	}
	winnings := storedGame.GetWagerCoin()
	if storedGame.MoveCount == 0 {
		panic(types.ErrNothingToPay.Error())
	} else if 1 < storedGame.MoveCount {
		winnings = winnings.Add(winnings)
	}
	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winnerAddress, sdk.NewCoins(winnings))
	if err != nil {
		panic(fmt.Sprintf(types.ErrCannotPayWinnings.Error(), err.Error()))
	}
}

func (k *Keeper) MustRefundWager(ctx sdk.Context, storedGame *types.StoredGame) {
	if storedGame.MoveCount == 1 {
		// Refund
		black, err := storedGame.GetBlackAddress()
		if err != nil {
			panic(err.Error())
		}
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, black, sdk.NewCoins(storedGame.GetWagerCoin()))
		if err != nil {
			panic(fmt.Sprintf(types.ErrCannotRefundWager.Error(), err.Error()))
		}
	} else if storedGame.MoveCount == 0 {
		// Do nothing
	} else {
		// TODO Implement a draw mechanism.
		panic(fmt.Sprintf(types.ErrNotInRefundState.Error(), storedGame.MoveCount))
	}
}
