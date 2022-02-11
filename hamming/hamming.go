package hamming

import (
	"fmt"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := fmt.Errorf("length is not the same")
		return 0, err
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
