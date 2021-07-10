package keeper

import (
	"github.com/hleb-albau/registry/x/registry/types"
)

var _ types.QueryServer = Keeper{}
