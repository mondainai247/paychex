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

func (k Keeper) EmployeeAll(goCtx context.Context, req *types.QueryAllEmployeeRequest) (*types.QueryAllEmployeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var employees []types.Employee
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	employeeStore := prefix.NewStore(store, types.KeyPrefix(types.EmployeeKey))

	pageRes, err := query.Paginate(employeeStore, req.Pagination, func(key []byte, value []byte) error {
		var employee types.Employee
		if err := k.cdc.Unmarshal(value, &employee); err != nil {
			return err
		}

		employees = append(employees, employee)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEmployeeResponse{Employee: employees, Pagination: pageRes}, nil
}

func (k Keeper) Employee(goCtx context.Context, req *types.QueryGetEmployeeRequest) (*types.QueryGetEmployeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	employee, found := k.GetEmployee(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEmployeeResponse{Employee: employee}, nil
}
