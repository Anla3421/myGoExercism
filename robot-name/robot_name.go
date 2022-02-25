package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var nameList = map[string]int{}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	format := "%s%s%.3v"
	rand.Seed(time.Now().UnixNano())
	for {
		s1 := string(byte(rand.Intn(25) + 65)) // random A~Z
		s2 := string(byte(rand.Intn(25) + 65))
		s3 := rand.Intn(999)
		r.name = fmt.Sprintf(format, s1, s2, s3)
		if _, ok := nameList[r.name]; !ok {
			nameList[r.name]++
			return r.name, nil
		}
	}

}

func (r *Robot) Reset() {
	r.name = ""
	nameList = map[string]int{}
}
