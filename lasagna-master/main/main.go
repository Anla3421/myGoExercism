package main

import (
	"fmt"
	"lasagna"
)

func main() {
	aaa := lasagna.ScaleRecipe([]float64{0.6, 300, 1, 0.5, 50, 0.1, 100}, 3)
	fmt.Println(aaa)
}
