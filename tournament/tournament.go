package tournament

import (
	"fmt"
	"io"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	p := make([]byte, 1024)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Printf("n: %d, err: %v\n", n, err)
			break
		}
		// fmt.Println(n, string(p[:n]))
	}
	fmt.Println("this is P", string(p))
	tempSlice := strings.Split(string(p), "\n")
	fmt.Println("tempSlice 0", tempSlice[0])
	return nil
}
