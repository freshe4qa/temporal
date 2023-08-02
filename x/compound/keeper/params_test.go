package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "temporal/testutil/keeper"
	"temporal/x/compound/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CompoundKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
