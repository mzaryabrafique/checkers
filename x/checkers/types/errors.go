package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidSigner = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
)

var (
	ErrInvalidBlack     = errors.Register(ModuleName, 1101, "black address is invalid: %s")
	ErrInvalidRed       = errors.Register(ModuleName, 1102, "red address is invalid: %s")
	ErrGameNotParseable = errors.Register(ModuleName, 1103, "game cannot be parsed")
)
