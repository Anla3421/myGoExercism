package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length is not the same")
	}
	var count int
	newA := strings.Split(a, "")
	newB := strings.Split(b, "")
	for i := 0; i < len(a); i++ {
		if newA[i] != newB[i] {
			count += 1
		}
	}
	return count, nil
}
