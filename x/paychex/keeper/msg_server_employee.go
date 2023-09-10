package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"paychex/x/paychex/types"
)

func (k msgServer) CreateEmployee(goCtx context.Context, msg *types.MsgCreateEmployee) (*types.MsgCreateEmployeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var employee = types.Employee{
		Creator: msg.Creator,
		Name:    msg.Name,
		Role:    msg.Role,
	}

	id := k.AppendEmployee(
		ctx,
		employee,
	)

	return &types.MsgCreateEmployeeResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateEmployee(goCtx context.Context, msg *types.MsgUpdateEmployee) (*types.MsgUpdateEmployeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var employee = types.Employee{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Role:    msg.Role,
	}

	// Checks that the element exists
	val, found := k.GetEmployee(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetEmployee(ctx, employee)

	return &types.MsgUpdateEmployeeResponse{}, nil
}

func (k msgServer) DeleteEmployee(goCtx context.Context, msg *types.MsgDeleteEmployee) (*types.MsgDeleteEmployeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetEmployee(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEmployee(ctx, msg.Id)

	return &types.MsgDeleteEmployeeResponse{}, nil
}
