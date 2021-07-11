package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hleb-albau/registry/x/registry/types"
)

func TestChainMsgRegister(t *testing.T) {
	// positive path
	srv, k, ctx := setupMsgServer(t)
	sdkContext := sdk.UnwrapSDKContext(ctx)
	owner := NewAccAddress()
	for i := 0; i < 5; i++ {
		_, err := srv.RegisterChain(ctx, &types.MsgRegisterChain{Owner: owner, ChainID: strconv.Itoa(i)})
		require.NoError(t, err)
		assert.Equal(t, uint64(i+1), k.GetChainCount(sdkContext))
	}
	// can't create chain with already occupied id
	_, err := srv.RegisterChain(ctx, &types.MsgRegisterChain{Owner: owner, ChainID: "some-id"})
	require.NoError(t, err)
	_, err = srv.RegisterChain(ctx, &types.MsgRegisterChain{Owner: owner, ChainID: "some-id"})
	require.Error(t, err)

	// can't create chain with not allowed characters in chainID
	_, err = srv.RegisterChain(ctx, &types.MsgRegisterChain{Owner: "A", ChainID: "chain{x}"})
	require.Error(t, err)
}

func TestChainMsgServerUpdate(t *testing.T) {
	owner := NewAccAddress()

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateChain
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateChain{Owner: owner, ChainID: "1"},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateChain{Owner: NewAccAddress(), ChainID: "2"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, _, ctx := setupMsgServer(t)
			_, err := srv.RegisterChain(ctx, &types.MsgRegisterChain{Owner: owner, ChainID: tc.request.ChainID})
			require.NoError(t, err)

			_, err = srv.UpdateChain(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
