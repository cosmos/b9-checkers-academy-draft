package types_test

import (
	"testing"

	"github.com/b9lab/checkers/testutil/sample"
	"github.com/b9lab/checkers/x/checkers/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGame_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateGame
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgCreateGame{
				Creator: "invalid_address",
				Black:   sample.AccAddress(),
				Red:     sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid black address",
			msg: types.MsgCreateGame{
				Creator: sample.AccAddress(),
				Black:   "invalid_address",
				Red:     sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid red address",
			msg: types.MsgCreateGame{
				Creator: sample.AccAddress(),
				Black:   sample.AccAddress(),
				Red:     "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid addresses",
			msg: types.MsgCreateGame{
				Creator: sample.AccAddress(),
				Black:   sample.AccAddress(),
				Red:     sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
