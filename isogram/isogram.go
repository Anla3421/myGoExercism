package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	counts := map[string]int{}
	temp := strings.Split(strings.ToLower(word), "")

	for k, v := range temp {
		if strings.IndexFunc(v, runeCheck) != -1 {
			temp = append(temp[:k], temp[k+1:]...)
		}
	}
	for _, v := range temp {
		counts[v]++
		if counts[v] > 1 {
			return false
		}
	}
	return true
}

func runeCheck(c rune) bool {
	return c < 'a' || c > 'z'
}
