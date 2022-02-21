package main

import (
	"fmt"
	"luhn"
)

func main() {
	// id := "4539 3195 0343 6467"
	id := "059"
	result := luhn.Valid(id)
	fmt.Println("result", result)
}
