package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cbank module sentinel errors
var (
	ErrInvalidSignature   = sdkerrors.Register(ModuleName, 1100, "failed to verify Schnorr signature")
	ErrCommitmentNotFound = sdkerrors.Register(ModuleName, 1101, "commitment not found")
	ErrCommitmentExist    = sdkerrors.Register(ModuleName, 1101, "commitment already exist")
	ErrInvalidRangeProof  = sdkerrors.Register(ModuleName, 1102, "failed to verify range-proof")
	ErrFailedRandom       = sdkerrors.Register(ModuleName, 1103, "failed to generate secure random")
)
