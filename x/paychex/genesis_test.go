package paychex_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "paychex/testutil/keeper"
	"paychex/testutil/nullify"
	"paychex/x/paychex"
	"paychex/x/paychex/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EmployeeList: []types.Employee{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EmployeeCount: 2,
		QueryResponseList: []types.QueryResponse{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		QueryResponseCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PaychexKeeper(t)
	paychex.InitGenesis(ctx, *k, genesisState)
	got := paychex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EmployeeList, got.EmployeeList)
	require.Equal(t, genesisState.EmployeeCount, got.EmployeeCount)
	require.ElementsMatch(t, genesisState.QueryResponseList, got.QueryResponseList)
	require.Equal(t, genesisState.QueryResponseCount, got.QueryResponseCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
