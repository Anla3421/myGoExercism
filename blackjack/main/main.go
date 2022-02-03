package main

import (
	"blackjack"
	"fmt"
)

func main() {
	text := blackjack.LargeHand(false, 20)
	fmt.Println(text)
}
