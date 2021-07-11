package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterChain{}

func NewMsgRegisterChain(chainID string, owner string) *MsgRegisterChain {
	return &MsgRegisterChain{
		ChainID: chainID,
		Owner:   owner,
	}
}

func (msg *MsgRegisterChain) Route() string {
	return RouterKey
}

func (msg *MsgRegisterChain) Type() string {
	return "RegisterChain"
}

func (msg *MsgRegisterChain) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgRegisterChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !validateChainID(msg.ChainID) {
		return sdkerrors.Wrap(ErrInvalidChain, "chainID contains not allowed characters")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateChain{}

func NewMsgUpdateChain(chainID string, owner string) *MsgUpdateChain {
	return &MsgUpdateChain{
		ChainID: chainID,
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
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func validateChainID(chainID string) bool {
	if len(chainID) == 0 {
		return false
	}
	for _, r := range chainID {
		if !validChainIDChar(r) {
			return false
		}
	}
	return true
}

// only hyphen and alphanumeric character allowed
func validChainIDChar(c rune) bool {
	return c == '-' || ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}
