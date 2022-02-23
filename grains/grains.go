package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("input<=0")
	}
	result := math.Pow(2, float64(number-1))
	return uint64(result), nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		temp, _ := Square(i)
		sum += temp

	}
	return sum
}
