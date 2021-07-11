package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	utils "github.com/hleb-albau/registry/testutil"
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
	owner := utils.NewAccAddress()
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
	owner := utils.NewAccAddress()

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
			request: &types.MsgUpdateChain{Owner: utils.NewAccAddress(), ChainID: "2"},
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

func TestChainMsgTransferOwnership(t *testing.T) {
	// positive
	newOwner := utils.NewAccAddress()
	srv, _, ctx := setupMsgServer(t)
	chain, err := srv.RegisterChain(ctx, &types.MsgRegisterChain{Owner: utils.NewAccAddress(), ChainID: "0"})
	require.NoError(t, err)
	_, err = srv.UpdateChain(ctx, &types.MsgUpdateChain{Owner: newOwner, ChainID: "0"})
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	_, err = srv.TransferChainOwnership(ctx, &types.MsgTransferChainOwnership{Owner: chain.Owner, NewOwner: newOwner, ChainID: "0"})
	require.NoError(t, err)
	_, err = srv.UpdateChain(ctx, &types.MsgUpdateChain{Owner: newOwner, ChainID: "0"})
	require.NoError(t, err)

	// can't transfer chain, that you don't own
	_, err = srv.RegisterChain(ctx, &types.MsgRegisterChain{Owner: utils.NewAccAddress(), ChainID: "1"})
	require.NoError(t, err)
	msg := types.MsgTransferChainOwnership{Owner: utils.NewAccAddress(), NewOwner: utils.NewAccAddress(), ChainID: "1"}
	_, err = srv.TransferChainOwnership(ctx, &msg)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)

	// can't transfer non-existing chain
	msg = types.MsgTransferChainOwnership{Owner: utils.NewAccAddress(), NewOwner: utils.NewAccAddress(), ChainID: "-1"}
	_, err = srv.TransferChainOwnership(ctx, &msg)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}
