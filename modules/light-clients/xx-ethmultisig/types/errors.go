package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidProof          = sdkerrors.Register(ModuleName, 1, "invalid Multisig proof")
	ErrInvalidSignatureCount = sdkerrors.Register(ModuleName, 2, "invalid signature count")
)
