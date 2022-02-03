package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	output := "Welcome to my party, " + name + "!"
	return output
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	output := "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
	return output
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	text := `Welcome to my party, %s!
You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.
You will be sitting next to %s.`
	output := fmt.Sprintf(text, name, table, direction, distance, neighbor)
	fmt.Println(output)
	return output
}
