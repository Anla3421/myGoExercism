package main

import (
	"fmt"
	"purchase"
)

func main() {
	text := purchase.ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
	fmt.Println(text)
}
