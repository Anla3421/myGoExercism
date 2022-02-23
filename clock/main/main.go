package main

import (
	"clock"
	"fmt"
)

func main() {
	// number := 0
	// c := clock.Clock{}
	result := clock.New(1, 10)
	result2 := clock.New(0, 50).Add(20)
	fmt.Println("result", result)
	fmt.Println("result2", result2)
	test := result == result2
	fmt.Println("test", test)
}
