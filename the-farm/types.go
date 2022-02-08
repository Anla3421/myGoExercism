package thefarm

import (
	"errors"
	"fmt"
)

// This file contains types used in the exercise but should not be modified.

// WeightFodder returns the amount of available fodder.
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// ErrScaleMalfunction indicates an error with the scale.
var ErrScaleMalfunction = errors.New("sensor error")
var ErrNegativeFodder = fmt.Errorf("negative fodder")
var ErrNonScale = fmt.Errorf("non-scale error")
var ErrDivisionByZero = fmt.Errorf("division by zero")
