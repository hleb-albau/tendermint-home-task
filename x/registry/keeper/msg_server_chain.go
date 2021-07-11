package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hleb-albau/registry/x/registry/types"
)

func (k msgServer) RegisterChain(goCtx context.Context, msg *types.MsgRegisterChain) (*types.Chain, error) {
	err := msg.ValidateBasic()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if k.HasChain(ctx, msg.ChainID) {
		errMsg := fmt.Sprintf("chain with chainID %s already exist", msg.ChainID)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errMsg)
	}

	var chain = types.Chain{
		ChainID: msg.ChainID,
		Owner:   msg.Owner,
	}

	k.SetChain(ctx, chain)
	chainsCount := k.GetChainCount(ctx)
	k.SetChainCount(ctx, chainsCount+1)
	return &chain, nil
}

func (k msgServer) UpdateChain(goCtx context.Context, msg *types.MsgUpdateChain) (*types.Chain, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var chain = types.Chain{
		ChainID: msg.ChainID,
		Owner:   msg.Owner,
	}

	if !k.HasChain(ctx, msg.ChainID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.ChainID))
	}
	// Checks if the msg sender is the same as the current owner
	if msg.Owner != k.GetChainOwner(ctx, msg.ChainID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetChain(ctx, chain)

	return &chain, nil
}
