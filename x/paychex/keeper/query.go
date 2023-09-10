package keeper

import (
	"paychex/x/paychex/types"
)

var _ types.QueryServer = Keeper{}
