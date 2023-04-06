package tournament

import (
	"fmt"
	"io"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	p := make([]byte, 1024)
	fileTable := map[string][]int{
		"Devastating Donkeys":     {0, 0, 0},
		"Allegoric Alaskians":     {0, 0, 0},
		"Blithering Badgers":      {0, 0, 0},
		"Courageous Californians": {0, 0, 0},
	}

	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			fmt.Printf("n: %d, err: %v\n", n, err)
			break
		}
		// fmt.Println(n, string(p[:n]))
	}
	// fmt.Println("this is P", string(p))
	tempSlice := strings.Split(string(p), "\n")
	for _, v := range tempSlice {
		mp := 1
		w := 1
		d := 1
		tempSlice2 := strings.Split(v, ";")
		if len(tempSlice2) == 3 {
			switch tempSlice2[2] {
			case "win":
				fileTable[tempSlice2[0]][0] = fileTable[tempSlice2[0]][0] + mp
				fileTable[tempSlice2[0]][1] = fileTable[tempSlice2[0]][1] + w

				fileTable[tempSlice2[1]][0] = fileTable[tempSlice2[1]][0] + mp
			case "loss":
				fileTable[tempSlice2[0]][0] = fileTable[tempSlice2[0]][0] + mp

				fileTable[tempSlice2[1]][0] = fileTable[tempSlice2[1]][0] + mp
				fileTable[tempSlice2[1]][1] = fileTable[tempSlice2[1]][1] + w
			case "draw":
				fileTable[tempSlice2[0]][0] = fileTable[tempSlice2[0]][0] + mp
				fileTable[tempSlice2[0]][2] = fileTable[tempSlice2[0]][2] + d

				fileTable[tempSlice2[1]][0] = fileTable[tempSlice2[1]][0] + mp
				fileTable[tempSlice2[1]][2] = fileTable[tempSlice2[1]][2] + d
			}
		}
	}

	teamDD := []int{
		fileTable["Devastating Donkeys"][0], fileTable["Devastating Donkeys"][1], fileTable["Devastating Donkeys"][2], fileTable["Devastating Donkeys"][0] - fileTable["Devastating Donkeys"][1] - fileTable["Devastating Donkeys"][2], fileTable["Devastating Donkeys"][1]*3 + fileTable["Devastating Donkeys"][2],
	}
	teamAA := []int{
		fileTable["Allegoric Alaskians"][0], fileTable["Allegoric Alaskians"][1], fileTable["Allegoric Alaskians"][2], fileTable["Allegoric Alaskians"][0] - fileTable["Allegoric Alaskians"][1] - fileTable["Allegoric Alaskians"][2], fileTable["Allegoric Alaskians"][1]*3 + fileTable["Allegoric Alaskians"][2],
	}
	teamBB := []int{
		fileTable["Blithering Badgers"][0], fileTable["Blithering Badgers"][1], fileTable["Blithering Badgers"][2], fileTable["Blithering Badgers"][0] - fileTable["Blithering Badgers"][1] - fileTable["Blithering Badgers"][2], fileTable["Blithering Badgers"][1]*3 + fileTable["Blithering Badgers"][2],
	}
	teamCC := []int{
		fileTable["Courageous Californians"][0], fileTable["Courageous Californians"][1], fileTable["Courageous Californians"][2], fileTable["Courageous Californians"][0] - fileTable["Courageous Californians"][1] - fileTable["Courageous Californians"][2], fileTable["Courageous Californians"][1]*3 + fileTable["Courageous Californians"][2],
	}
	result := fmt.Sprintf(`
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  %v |  %v |  %v |  %v |  %v
Allegoric Alaskians            |  %v |  %v |  %v |  %v |  %v
Blithering Badgers             |  %v |  %v |  %v |  %v |  %v
Courageous Californians        |  %v |  %v |  %v |  %v |  %v
`, teamDD[0], teamDD[1], teamDD[2], teamDD[3], teamDD[4],
		teamAA[0], teamAA[1], teamAA[2], teamAA[3], teamAA[4],
		teamBB[0], teamBB[1], teamBB[2], teamBB[3], teamBB[4],
		teamCC[0], teamCC[1], teamCC[2], teamCC[3], teamCC[4],
	)
	fmt.Println("result", result)
	n, err := writer.Write([]byte(result))
	if err == io.EOF {
		fmt.Printf("n: %d, err: %v\n", n, err)
	}
	return nil
}
