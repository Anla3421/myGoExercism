package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var sum int
	newID := []int{}
	temp := strings.Split(strings.ReplaceAll(id, " ", ""), "")
	if len(temp) < 2 {
		return false
	}
	for i := len(temp) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(temp[i])
		if err != nil {
			return false
		}
		newID = append(newID, num)
	}
	for k, v := range newID {
		if k%2 == 1 {
			if v >= 5 {
				v = 2*v%10 + 1
			} else {
				v = 2 * v % 10
			}
		}
		sum += v
	}
	if sum%10 != 0 {
		return false
	}
	return true
}
