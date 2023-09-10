package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "paychex/testutil/keeper"
	"paychex/testutil/nullify"
	"paychex/x/paychex/keeper"
	"paychex/x/paychex/types"
)

func createNEmployee(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Employee {
	items := make([]types.Employee, n)
	for i := range items {
		items[i].Id = keeper.AppendEmployee(ctx, items[i])
	}
	return items
}

func TestEmployeeGet(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNEmployee(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEmployee(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEmployeeRemove(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNEmployee(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEmployee(ctx, item.Id)
		_, found := keeper.GetEmployee(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEmployeeGetAll(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNEmployee(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEmployee(ctx)),
	)
}

func TestEmployeeCount(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNEmployee(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEmployeeCount(ctx))
}
