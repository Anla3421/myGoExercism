package thefarm

import (
	"errors"
	"fmt"
)

//
type animal interface {
	Bark() int
}

// SillyNephewError alerts when given negative cows.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

func (e *SillyNephewError) Bark() int {
	return 123
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error, animal) {
	fodder, err := weightFodder.FodderAmount()
	fmt.Println(&SillyNephewError{cows: cows})
	fmt.Println(SillyNephewError{cows: cows})
	switch {
	case err == ErrScaleMalfunction:
		fodder *= 2
	case err != nil:
		return 0, err
	case fodder < 0:
		return 0, errors.New("Negative fodder")
	case cows == 0:
		return 0, errors.New("Division by zero")
	case cows < 0:
		return 0, &SillyNephewError{cows: cows}, &SillyNephewError{cows: cows}
	}
	return fodder / float64(cows), nil
}
