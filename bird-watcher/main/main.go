package main

import "fmt"

func main() {
	// birdwatcher.BirdsInWeek([]int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0, 8, 0}, 1)
	aaa, bbb := SumAndMultiplyThenMinus(1, 2, 3)
	fmt.Println(aaa, bbb)
}

func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
	sum, mult = a+b, a*b //3 2
	sum += c             //3
	mult -= c            //3
	return
}
