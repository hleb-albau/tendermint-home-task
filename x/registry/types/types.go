package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidChain = sdkerrors.Register(ModuleName, 1, "invalid chain")
)
