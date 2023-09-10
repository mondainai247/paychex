package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"paychex/x/paychex/types"
)

func (k msgServer) SendPayroll(goCtx context.Context, msg *types.MsgSendPayroll) (*types.MsgSendPayrollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendPayrollResponse{}, nil
}
