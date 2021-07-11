package keeper

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hleb-albau/registry/x/registry/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, Keeper, context.Context) {
	keeper, ctx := setupKeeper(t)
	return NewMsgServerImpl(*keeper), *keeper, sdk.WrapSDKContext(ctx)
}
