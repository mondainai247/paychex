package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"paychex/testutil/sample"
)

func TestMsgCreateQueryResponse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateQueryResponse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateQueryResponse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateQueryResponse{
				Creator: sample.AccAddress(),
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

func TestMsgUpdateQueryResponse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateQueryResponse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateQueryResponse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateQueryResponse{
				Creator: sample.AccAddress(),
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

func TestMsgDeleteQueryResponse_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteQueryResponse
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteQueryResponse{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteQueryResponse{
				Creator: sample.AccAddress(),
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
