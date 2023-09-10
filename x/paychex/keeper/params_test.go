package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "paychex/testutil/keeper"
	"paychex/x/paychex/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PaychexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
