package main

import (
	"fmt"
	"tree"
)

type Record struct {
	ID     int
	Parent int
}

func new() []tree.Record {
	return []tree.Record{
		{ID: 0},
		{ID: 1, Parent: 0},
		{ID: 2, Parent: 0}}
}

func main() {
	records := new()
	result, err := tree.Build(records)
	fmt.Println("result", result, err)
}
