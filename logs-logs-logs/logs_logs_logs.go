package logs

import (
	"strings"
)

var (
	recommendation int32 = 0x2757
	search         int32 = 0x1F50D
	weather        int32 = 0x2600
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	changeToString := make(map[string]string)
	changeToString["recommendation"] = string(recommendation)
	changeToString["search"] = string(search)
	changeToString["weather"] = string(weather)
	for s, needToCompare := range changeToString {
		if strings.Contains(log, needToCompare) == true {
			return s
		}
	}
	// for pos, char := range "안녕히😀, hello ： A Α α π " {
	// 	fmt.Printf("character %c starts at byte position %d\n", char, pos)
	// }
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	oldString := string(oldRune)
	newString := string(newRune)
	output := strings.ReplaceAll(log, oldString, newString)
	return output
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	length := len([]rune(log))
	if length <= limit {
		return true
	}
	return false
}
