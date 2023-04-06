package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"tournament"
)

func main() {
	var myWriter io.Writer = os.Stdout
	var input io.Reader
	input = strings.NewReader(`Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskians;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskians;Courageous Californians;win
`)

	// 	input = strings.NewReader(`
	// Allegoric Alaskians;Blithering Badgers;win
	// Allegoric Alaskians;Courageous Californians;win
	// `)

	result := tournament.Tally(input, myWriter)
	fmt.Println("result", result)
}
