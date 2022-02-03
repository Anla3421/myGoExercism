// package main

// import "logs"

// func main() {
// 	// aaa := 0x1F50D
// 	// fmt.Printf("%c", aaa)
// 	logs.WithinLimit("❗ duct", 9)

// }

package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	b := "这是个测试"
	len1 := len([]rune(b))
	len2 := bytes.Count([]byte(b), nil) - 1
	len3 := strings.Count(b, "") - 1
	len4 := utf8.RuneCountInString(b)
	fmt.Println(len1)
	fmt.Println(len2)
	fmt.Println(len3)
	fmt.Println(len4)

	for pos, char := range "안녕히😀, hello ： A Α α π " {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	rune := 0x1F9E0
	fmt.Printf("%c", rune)
}
