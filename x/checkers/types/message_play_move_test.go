package types_test

import (
	"testing"

	"github.com/b9lab/checkers/testutil/sample"
	"github.com/b9lab/checkers/x/checkers/rules"
	"github.com/b9lab/checkers/x/checkers/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPlayMove_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgPlayMove
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgPlayMove{
				Creator:   "invalid_address",
				GameIndex: "5",
				FromX:     0,
				FromY:     5,
				ToX:       1,
				ToY:       4,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid game index",
			msg: types.MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "invalid_index",
				FromX:     0,
				FromY:     5,
				ToX:       1,
				ToY:       4,
			},
			err: types.ErrInvalidGameIndex,
		},
		{
			name: "invalid fromX too high",
			msg: types.MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "5",
				FromX:     rules.BOARD_DIM,
				FromY:     5,
				ToX:       1,
				ToY:       4,
			},
			err: types.ErrInvalidPositionIndex,
		},
		{
			name: "invalid fromY too high",
			msg: types.MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "5",
				FromX:     0,
				FromY:     rules.BOARD_DIM,
				ToX:       1,
				ToY:       4,
			},
			err: types.ErrInvalidPositionIndex,
		},
		{
			name: "invalid toX too high",
			msg: types.MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "5",
				FromX:     0,
				FromY:     5,
				ToX:       rules.BOARD_DIM,
				ToY:       4,
			},
			err: types.ErrInvalidPositionIndex,
		},
		{
			name: "invalid toY too high",
			msg: types.MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "5",
				FromX:     0,
				FromY:     5,
				ToX:       1,
				ToY:       rules.BOARD_DIM,
			},
			err: types.ErrInvalidPositionIndex,
		},
		{
			name: "invalid no move",
			msg: types.MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "5",
				FromX:     0,
				FromY:     5,
				ToX:       0,
				ToY:       5,
			},
			err: types.ErrMoveAbsent,
		},
		{
			name: "valid address",
			msg: types.MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "5",
				FromX:     0,
				FromY:     5,
				ToX:       1,
				ToY:       4,
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
