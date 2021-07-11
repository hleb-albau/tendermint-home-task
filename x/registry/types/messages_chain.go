package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateChain{}

func NewMsgCreateChain(creator string, chainId string, owner string) *MsgCreateChain {
	return &MsgCreateChain{
		Creator: creator,
		ChainId: chainId,
		Owner:   owner,
	}
}

func (msg *MsgCreateChain) Route() string {
	return RouterKey
}

func (msg *MsgCreateChain) Type() string {
	return "CreateChain"
}

func (msg *MsgCreateChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateChain{}

func NewMsgUpdateChain(creator string, id uint64, chainId string, owner string) *MsgUpdateChain {
	return &MsgUpdateChain{
		Id:      id,
		Creator: creator,
		ChainId: chainId,
		Owner:   owner,
	}
}

func (msg *MsgUpdateChain) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChain) Type() string {
	return "UpdateChain"
}

func (msg *MsgUpdateChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteChain{}

func NewMsgDeleteChain(creator string, id uint64) *MsgDeleteChain {
	return &MsgDeleteChain{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteChain) Route() string {
	return RouterKey
}

func (msg *MsgDeleteChain) Type() string {
	return "DeleteChain"
}

func (msg *MsgDeleteChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
