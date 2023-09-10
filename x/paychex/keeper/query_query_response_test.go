package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "paychex/testutil/keeper"
	"paychex/testutil/nullify"
	"paychex/x/paychex/types"
)

func TestQueryResponseQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNQueryResponse(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetQueryResponseRequest
		response *types.QueryGetQueryResponseResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetQueryResponseRequest{Id: msgs[0].Id},
			response: &types.QueryGetQueryResponseResponse{QueryResponse: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetQueryResponseRequest{Id: msgs[1].Id},
			response: &types.QueryGetQueryResponseResponse{QueryResponse: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetQueryResponseRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.QueryResponse(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestQueryResponseQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PaychexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNQueryResponse(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllQueryResponseRequest {
		return &types.QueryAllQueryResponseRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.QueryResponseAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.QueryResponse), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.QueryResponse),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.QueryResponseAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.QueryResponse), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.QueryResponse),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.QueryResponseAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.QueryResponse),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.QueryResponseAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
