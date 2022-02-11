package main

import (
	"fmt"
	"hamming"
)

func main() {
	a := "asdfg"
	b := "asdvb"
	result, err := hamming.Distance(a, b)
	fmt.Println("result", result)
	fmt.Println("err", err)
}
