package paychex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"paychex/x/paychex/keeper"
	"paychex/x/paychex/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the employee
	for _, elem := range genState.EmployeeList {
		k.SetEmployee(ctx, elem)
	}

	// Set employee count
	k.SetEmployeeCount(ctx, genState.EmployeeCount)
	// Set all the queryResponse
	for _, elem := range genState.QueryResponseList {
		k.SetQueryResponse(ctx, elem)
	}

	// Set queryResponse count
	k.SetQueryResponseCount(ctx, genState.QueryResponseCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EmployeeList = k.GetAllEmployee(ctx)
	genesis.EmployeeCount = k.GetEmployeeCount(ctx)
	genesis.QueryResponseList = k.GetAllQueryResponse(ctx)
	genesis.QueryResponseCount = k.GetQueryResponseCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
