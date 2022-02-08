package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	output := fmt.Sprintf("This is the number %.1f", f)
	return output
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	numGet := float64(nb.Number())
	output := fmt.Sprintf("This is a box containing the number %.1f", numGet)
	return output
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch text := fnb.(type) {
	case FancyNumber:
		output, err := strconv.Atoi(text.Value())
		if err != nil {
			fmt.Println(err)
			panic(0)
		}
		return output
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	temp := ExtractFancyNumber(fnb)
	output := fmt.Sprintf("This is a fancy box containing the number %.1f", float64(temp))
	return output
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch temp := i.(type) {
	case int:
		intTemp := float64(temp)
		output := DescribeNumber(intTemp)
		return output
	case float64:
		output := DescribeNumber(temp)
		return output
	case NumberBox:
		output := DescribeNumberBox(temp)
		return output
	case FancyNumberBox:
		output := DescribeFancyNumberBox(temp)
		return output
	default:
	}
	return "Return to sender"
}
