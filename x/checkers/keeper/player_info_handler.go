package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xavierlepretre/checkers/x/checkers/rules"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

func mustAddDeltaGameResultToPlayer(
	k *Keeper,
	ctx sdk.Context,
	player sdk.AccAddress,
	wonDelta uint64,
	lostDelta uint64,
	forfeitedDelta uint64,
) {
	playerInfo, found := k.GetPlayerInfo(ctx, player.String())
	if !found {
		playerInfo = types.PlayerInfo{
			Index:          player.String(),
			WonCount:       0,
			LostCount:      0,
			ForfeitedCount: 0,
		}
	}
	playerInfo.WonCount += wonDelta
	playerInfo.LostCount += lostDelta
	playerInfo.ForfeitedCount += forfeitedDelta
	k.SetPlayerInfo(ctx, playerInfo)
}

func (k *Keeper) MustAddWonGameResultToPlayer(ctx sdk.Context, player sdk.AccAddress) {
	mustAddDeltaGameResultToPlayer(k, ctx, player, 1, 0, 0)
}

func (k *Keeper) MustAddLostGameResultToPlayer(ctx sdk.Context, player sdk.AccAddress) {
	mustAddDeltaGameResultToPlayer(k, ctx, player, 0, 1, 0)
}

func (k *Keeper) MustAddForfeitedGameResultToPlayer(ctx sdk.Context, player sdk.AccAddress) {
	mustAddDeltaGameResultToPlayer(k, ctx, player, 0, 0, 1)
}

func getWinnerAndLoserAddresses(storedGame *types.StoredGame) (winnerAddress sdk.AccAddress, loserAddress sdk.AccAddress) {
	if storedGame.Winner == rules.NO_PLAYER.Color {
		panic(types.ErrThereIsNoWinner.Error())
	}
	redAddress, err := storedGame.GetRedAddress()
	if err != nil {
		panic(err.Error())
	}
	blackAddress, err := storedGame.GetBlackAddress()
	if err != nil {
		panic(err.Error())
	}
	if storedGame.Winner == rules.RED_PLAYER.Color {
		winnerAddress = redAddress
		loserAddress = blackAddress
	} else if storedGame.Winner == rules.BLACK_PLAYER.Color {
		winnerAddress = blackAddress
		loserAddress = redAddress
	} else {
		panic(fmt.Sprintf(types.ErrWinnerNotParseable.Error(), storedGame.Winner))
	}
	return winnerAddress, loserAddress
}

func (k *Keeper) MustRegisterPlayerWin(ctx sdk.Context, storedGame *types.StoredGame) {
	winnerAddress, loserAddress := getWinnerAndLoserAddresses(storedGame)
	k.MustAddWonGameResultToPlayer(ctx, winnerAddress)
	k.MustAddLostGameResultToPlayer(ctx, loserAddress)
}

func (k *Keeper) MustRegisterPlayerForfeit(ctx sdk.Context, storedGame *types.StoredGame) {
	winnerAddress, loserAddress := getWinnerAndLoserAddresses(storedGame)
	k.MustAddWonGameResultToPlayer(ctx, winnerAddress)
	k.MustAddForfeitedGameResultToPlayer(ctx, loserAddress)
}
