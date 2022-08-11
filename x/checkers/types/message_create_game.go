package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGame = "create_game"

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, black string, red string, wager uint64) *MsgCreateGame {
	return &MsgCreateGame{
		Creator: creator,
		Black:   black,
		Red:     red,
		Wager:   wager,
	}
}

func (msg *MsgCreateGame) Route() string {
	return RouterKey
}

func (msg *MsgCreateGame) Type() string {
	return TypeMsgCreateGame
}

func (msg *MsgCreateGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Black)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid black address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Red)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid red address (%s)", err)
	}
	return nil
}
