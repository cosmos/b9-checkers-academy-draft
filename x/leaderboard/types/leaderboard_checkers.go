package types

import (
	checkerstypes "github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func MakeCandidateFromPlayerInfo(playerInfo checkerstypes.PlayerInfo) (candidate Candidate, err error) {
	address, err := sdk.AccAddressFromBech32(playerInfo.Index)
	if err != nil {
		return candidate, sdkerrors.Wrapf(err, "Could not parse address from playerInfo %s", playerInfo.Index)
	}
	return Candidate{
		Address:  address,
		WonCount: playerInfo.WonCount,
	}, nil
}
