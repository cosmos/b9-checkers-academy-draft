package types

import (
	"strconv"

	"github.com/b9lab/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlayMove = "play_move"

var _ sdk.Msg = &MsgPlayMove{}

func NewMsgPlayMove(creator string, gameIndex string, fromX uint64, fromY uint64, toX uint64, toY uint64) *MsgPlayMove {
	return &MsgPlayMove{
		Creator:   creator,
		GameIndex: gameIndex,
		FromX:     fromX,
		FromY:     fromY,
		ToX:       toX,
		ToY:       toY,
	}
}

func (msg *MsgPlayMove) Route() string {
	return RouterKey
}

func (msg *MsgPlayMove) Type() string {
	return TypeMsgPlayMove
}

func (msg *MsgPlayMove) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPlayMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPlayMove) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	gameIndex, err := strconv.ParseInt(msg.GameIndex, 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidGameIndex, "not parseable (%s)", err)
	}
	if uint64(gameIndex) < DefaultIndex {
		return sdkerrors.Wrapf(ErrInvalidGameIndex, "number too low (%d)", gameIndex)
	}
	boardChecks := []struct {
		value uint64
		err   string
	}{
		{
			value: msg.FromX,
			err:   "fromX out of range (%d)",
		},
		{
			value: msg.ToX,
			err:   "toX out of range (%d)",
		},
		{
			value: msg.FromY,
			err:   "fromY out of range (%d)",
		},
		{
			value: msg.ToY,
			err:   "toY out of range (%d)",
		},
	}
	for _, situation := range boardChecks {
		if situation.value < 0 || rules.BOARD_DIM <= situation.value {
			return sdkerrors.Wrapf(ErrInvalidPositionIndex, situation.err, situation.value)
		}
	}
	if msg.FromX == msg.ToX && msg.FromY == msg.ToY {
		return sdkerrors.Wrapf(ErrMoveAbsent, "x (%d) and y (%d)", msg.FromX, msg.FromY)
	}
	return nil
}
