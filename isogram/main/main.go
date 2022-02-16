package main

import (
	"fmt"
	"isogram"
)

func main() {
	a := "six-year-old"
	result := isogram.IsIsogram(a)
	fmt.Println("result", result)
}

// func main() {

// 	fmt.Println(strings.IndexFunc("Hello, 世界", f))
// 	fmt.Println(strings.IndexFunc("Hello, world", f))
// 	fmt.Println(strings.IndexFunc("Helloworld", f))
// }
