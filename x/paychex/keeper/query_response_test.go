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

func createNQueryResponse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.QueryResponse {
	items := make([]types.QueryResponse, n)
	for i := range items {
		items[i].Id = keeper.AppendQueryResponse(ctx, items[i])
	}
	return items
}

func TestQueryResponseGet(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNQueryResponse(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetQueryResponse(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestQueryResponseRemove(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNQueryResponse(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveQueryResponse(ctx, item.Id)
		_, found := keeper.GetQueryResponse(ctx, item.Id)
		require.False(t, found)
	}
}

func TestQueryResponseGetAll(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNQueryResponse(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllQueryResponse(ctx)),
	)
}

func TestQueryResponseCount(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	items := createNQueryResponse(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetQueryResponseCount(ctx))
}
