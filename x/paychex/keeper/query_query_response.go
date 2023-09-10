package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"paychex/x/paychex/types"
)

func (k Keeper) QueryResponseAll(goCtx context.Context, req *types.QueryAllQueryResponseRequest) (*types.QueryAllQueryResponseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var queryResponses []types.QueryResponse
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	queryResponseStore := prefix.NewStore(store, types.KeyPrefix(types.QueryResponseKey))

	pageRes, err := query.Paginate(queryResponseStore, req.Pagination, func(key []byte, value []byte) error {
		var queryResponse types.QueryResponse
		if err := k.cdc.Unmarshal(value, &queryResponse); err != nil {
			return err
		}

		queryResponses = append(queryResponses, queryResponse)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllQueryResponseResponse{QueryResponse: queryResponses, Pagination: pageRes}, nil
}

func (k Keeper) QueryResponse(goCtx context.Context, req *types.QueryGetQueryResponseRequest) (*types.QueryGetQueryResponseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	queryResponse, found := k.GetQueryResponse(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetQueryResponseResponse{QueryResponse: queryResponse}, nil
}
