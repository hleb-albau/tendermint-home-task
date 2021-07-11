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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgTransferChainOwnership{}

func NewMsgTransferChainOwnership(chainID string, owner string, newOwner string) *MsgTransferChainOwnership {
	return &MsgTransferChainOwnership{
		ChainID:  chainID,
		Owner:    owner,
		NewOwner: newOwner,
	}
}

func (m *MsgTransferChainOwnership) Route() string {
	return RouterKey
}

func (m *MsgTransferChainOwnership) Type() string {
	return "TransferChainOwnership"
}

func (m *MsgTransferChainOwnership) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(m.NewOwner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new owner address (%s)", err)
	}
	return nil
}

func (m *MsgTransferChainOwnership) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgTransferChainOwnership) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
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
