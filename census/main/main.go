package main

import (
	"census"
)

// type rect struct {
// 	width, height int
// }

// func (r *rect) squareIt() {
// 	r.height = r.width
// }

func main() {
	c := census.NewResident("Matthew Sanabria", 29,
		map[string]string{})
	c.HasRequiredInfo()
}
