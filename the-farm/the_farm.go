package thefarm

import (
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows <= 0 {
		if cows == 0 {
			err := ErrDivisionByZero
			return 0, err
		} else {
			err := &SillyNephewError{cows: cows}
			return 0, err
		}
	}
	fodder, err := weightFodder.FodderAmount()
	if fodder < 0 {
		if err == nil {
			err = ErrNegativeFodder
		}
		if err != nil {
			if err.Error() == ErrNonScale.Error() {
				err = ErrNonScale
				return 0, err
			}
			err = ErrNegativeFodder
			return 0, err
		}
	}
	if err != nil {
		switch err.Error() {
		case ErrScaleMalfunction.Error():
			fodder = fodder * 2
			err = nil
		case ErrNonScale.Error():
			err := ErrNonScale
			return 0, err
		default:
			fmt.Println("here is default")
		}
	}
	output := fodder / float64(cows)
	return output, err
}
