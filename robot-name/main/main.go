package main

import (
	"fmt"
	"robotname"
)

func main() {
	r := robotname.Robot{}
	result, err := r.Name()

	fmt.Println("result", result, err)
}
