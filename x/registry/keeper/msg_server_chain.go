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
	if _, err := k.verifyOwnership(ctx, msg.ChainID, msg.Owner); err != nil {
		return nil, err
	}

	var chain = types.Chain{
		ChainID: msg.ChainID,
		Owner:   msg.Owner,
	}

	k.SetChain(ctx, chain)
	return &chain, nil
}

func (k msgServer) TransferChainOwnership(
	goCtx context.Context, msg *types.MsgTransferChainOwnership,
) (*types.Chain, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	chain, err := k.verifyOwnership(ctx, msg.ChainID, msg.Owner)
	if err != nil {
		return nil, err
	}
	chain.Owner = msg.NewOwner
	k.SetChain(ctx, *chain)
	err = ctx.EventManager().EmitTypedEvent(
		&types.EventChainOwnershipTransfer{
			ChainID:  msg.ChainID,
			Owner:    msg.Owner,
			NewOwner: msg.NewOwner,
		},
	)
	return chain, err
}

func (k msgServer) verifyOwnership(ctx sdk.Context, chainID string, owner string) (*types.Chain, error) {
	if !k.HasChain(ctx, chainID) {
		errMsg := fmt.Sprintf("chain with chainID '%s' doesn't exist", chainID)
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, errMsg)
	}
	chain := k.GetChain(ctx, chainID)
	if owner != chain.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	return &chain, nil
}
