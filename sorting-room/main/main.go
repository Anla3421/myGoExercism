package main

import "fmt"

func main() {
	var input interface{} = 12
	number := input.(float64)
	str, ok := input.(float64)
	fmt.Println("input:", input)
	fmt.Println("number:", number)
	fmt.Println("str:", str)
	fmt.Println("ok:", ok)
	switch v := input.(type) {
	case int:
		fmt.Printf("the integer %d", v)
	case string:
		fmt.Printf("the string %s", v)
	default:
		fmt.Println("some type we did not handle explicitly")
	}
}
