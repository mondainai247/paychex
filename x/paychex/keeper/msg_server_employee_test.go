package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"paychex/x/paychex/types"
)

func TestEmployeeMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateEmployee(ctx, &types.MsgCreateEmployee{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestEmployeeMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateEmployee
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateEmployee{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEmployee{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEmployee{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateEmployee(ctx, &types.MsgCreateEmployee{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateEmployee(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestEmployeeMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteEmployee
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteEmployee{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteEmployee{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteEmployee{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateEmployee(ctx, &types.MsgCreateEmployee{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteEmployee(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
