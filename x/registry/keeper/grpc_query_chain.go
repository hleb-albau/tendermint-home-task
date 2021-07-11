package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hleb-albau/registry/x/registry/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChainAll(c context.Context, req *types.QueryAllChainRequest) (*types.QueryAllChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var chains []*types.Chain
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	chainStore := prefix.NewStore(store, types.KeyPrefix(types.ChainKey))

	pageRes, err := query.Paginate(chainStore, req.Pagination, func(key []byte, value []byte) error {
		var chain types.Chain
		if err := k.cdc.UnmarshalBinaryBare(value, &chain); err != nil {
			return err
		}

		chains = append(chains, &chain)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChainResponse{Chain: chains, Pagination: pageRes}, nil
}

func (k Keeper) Chain(c context.Context, req *types.QueryGetChainRequest) (*types.QueryGetChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var chain types.Chain
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasChain(ctx, req.ChainID) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetChainKey(req.ChainID)), &chain)

	return &types.QueryGetChainResponse{Chain: &chain}, nil
}
