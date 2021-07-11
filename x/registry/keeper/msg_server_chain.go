package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hleb-albau/registry/x/registry/types"
)

func (k msgServer) CreateChain(goCtx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var chain = types.Chain{
		Creator: msg.Creator,
		ChainId: msg.ChainId,
		Owner:   msg.Owner,
	}

	id := k.AppendChain(
		ctx,
		chain,
	)

	return &types.MsgCreateChainResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateChain(goCtx context.Context, msg *types.MsgUpdateChain) (*types.MsgUpdateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var chain = types.Chain{
		Creator: msg.Creator,
		Id:      msg.Id,
		ChainId: msg.ChainId,
		Owner:   msg.Owner,
	}

	// Checks that the element exists
	if !k.HasChain(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetChainOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetChain(ctx, chain)

	return &types.MsgUpdateChainResponse{}, nil
}

func (k msgServer) DeleteChain(goCtx context.Context, msg *types.MsgDeleteChain) (*types.MsgDeleteChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasChain(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetChainOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveChain(ctx, msg.Id)

	return &types.MsgDeleteChainResponse{}, nil
}
