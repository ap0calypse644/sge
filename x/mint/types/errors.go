package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mint module sentinel errors
var (
	ErrMintDenomIsBlank = sdkerrors.Register(ModuleName, 1100, "mint denom cannot be blank")
)

// nolint
const (
	ErrTextInvalidParamType                   = "invalid parameter type: %T"
	ErrTextBlocksPerYearMustBePositive        = "blocks per year must be positive: %d"
	ErrTextExcludeAmountMustBePositive        = "exclude amount must be positive: %s"
	ErrTextPhasesShouldHaveValue              = "phases should have value: %v"
	ErrTextMintParamInflationShouldBePositive = "mint parameter Inflation should be positive, is %s"
	ErrTextYearCoefficientMustBePositive      = "year coeficient should be non-zero and positive value"
	ErrTextEndPhaseParamNotAllowed            = "adding phase with equal values with end phase is not allowed"
	ErrTextNilMinter                          = "stored minter should not have been nil"
)
