package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"paychex/x/paychex/types"
)

func (k msgServer) CreateQueryResponse(goCtx context.Context, msg *types.MsgCreateQueryResponse) (*types.MsgCreateQueryResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var queryResponse = types.QueryResponse{
		Creator: msg.Creator,
		Staff:   msg.Staff,
	}

	id := k.AppendQueryResponse(
		ctx,
		queryResponse,
	)

	return &types.MsgCreateQueryResponseResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateQueryResponse(goCtx context.Context, msg *types.MsgUpdateQueryResponse) (*types.MsgUpdateQueryResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var queryResponse = types.QueryResponse{
		Creator: msg.Creator,
		Id:      msg.Id,
		Staff:   msg.Staff,
	}

	// Checks that the element exists
	val, found := k.GetQueryResponse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetQueryResponse(ctx, queryResponse)

	return &types.MsgUpdateQueryResponseResponse{}, nil
}

func (k msgServer) DeleteQueryResponse(goCtx context.Context, msg *types.MsgDeleteQueryResponse) (*types.MsgDeleteQueryResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetQueryResponse(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveQueryResponse(ctx, msg.Id)

	return &types.MsgDeleteQueryResponseResponse{}, nil
}
