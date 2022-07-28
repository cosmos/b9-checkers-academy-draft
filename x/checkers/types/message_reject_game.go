package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectGame = "reject_game"

var _ sdk.Msg = &MsgRejectGame{}

func NewMsgRejectGame(creator string, gameIndex string) *MsgRejectGame {
	return &MsgRejectGame{
		Creator:   creator,
		GameIndex: gameIndex,
	}
}

func (msg *MsgRejectGame) Route() string {
	return RouterKey
}

func (msg *MsgRejectGame) Type() string {
	return TypeMsgRejectGame
}

func (msg *MsgRejectGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRejectGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
