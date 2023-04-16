package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
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
	}

	fileTable := Calculate(string(p))

	if fileTable["Devastating Donkeys"][0] == 0 || fileTable["Allegoric Alaskians"][0] == 0 || fileTable["Blithering Badgers"][0] == 0 || fileTable["Courageous Californians"][0] == 0 {
		return errors.New("teams error")
	}

	tempTable := map[int]int{}
	for _, v := range fileTable {
		tempTable[v[5]] = v[4]
	}
	fmt.Println("tempTable", tempTable)

	// 先將所有 key 存入 []string 變數
	keys := make([]string, 0, len(fileTable))
	for key := range fileTable {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return fileTable[keys[i]][len(fileTable[keys[i]])-2] > fileTable[keys[j]][len(fileTable[keys[j]])-2]
	})

	fmt.Println("keys", keys)
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
	// println("1111", teamDD, teamAA, teamBB, teamCC)

	result := fmt.Sprintf(`
	Team                           | MP |  W |  D |  L |  P
	Devastating Donkeys            |  %v |  %v |  %v |  %v |  %v
	Allegoric Alaskians            |  %v |  %v |  %v |  %v |  %v
	Blithering Badgers             |  %v |  %v |  %v |  %v |  %v
	Courageous Californians        |  %v |  %v |  %v |  %v |  %v
	`[1:], teamDD[0], teamDD[1], teamDD[2], teamDD[3], teamDD[4],
		teamAA[0], teamAA[1], teamAA[2], teamAA[3], teamAA[4],
		teamBB[0], teamBB[1], teamBB[2], teamBB[3], teamBB[4],
		teamCC[0], teamCC[1], teamCC[2], teamCC[3], teamCC[4],
	)
	// println("2222", result)

	n, err := writer.Write([]byte(result))
	if err == io.EOF {
		fmt.Printf("n: %d, err: %v\n", n, err)
	}
	return nil
}

func Calculate(p string) map[string][]int {
	// mp, win, draw, loss, point
	tempFileTable := map[string][]int{
		"Devastating Donkeys":     {0, 0, 0, 0, 0, 4},
		"Allegoric Alaskians":     {0, 0, 0, 0, 0, 1},
		"Blithering Badgers":      {0, 0, 0, 0, 0, 2},
		"Courageous Californians": {0, 0, 0, 0, 0, 3},
	}

	tempSlice := strings.Split(string(p), "\n")
	for _, v := range tempSlice {
		mp := 1
		w := 1
		d := 1
		tempSlice2 := strings.Split(v, ";")
		if len(tempSlice2) == 3 {
			switch tempSlice2[2] {
			case "win":
				tempFileTable[tempSlice2[0]][0] = tempFileTable[tempSlice2[0]][0] + mp
				tempFileTable[tempSlice2[0]][1] = tempFileTable[tempSlice2[0]][1] + w

				tempFileTable[tempSlice2[1]][0] = tempFileTable[tempSlice2[1]][0] + mp
			case "loss":
				tempFileTable[tempSlice2[0]][0] = tempFileTable[tempSlice2[0]][0] + mp

				tempFileTable[tempSlice2[1]][0] = tempFileTable[tempSlice2[1]][0] + mp
				tempFileTable[tempSlice2[1]][1] = tempFileTable[tempSlice2[1]][1] + w
			case "draw":
				tempFileTable[tempSlice2[0]][0] = tempFileTable[tempSlice2[0]][0] + mp
				tempFileTable[tempSlice2[0]][2] = tempFileTable[tempSlice2[0]][2] + d

				tempFileTable[tempSlice2[1]][0] = tempFileTable[tempSlice2[1]][0] + mp
				tempFileTable[tempSlice2[1]][2] = tempFileTable[tempSlice2[1]][2] + d
			}
		}
	}
	fileTable := map[string][]int{}
	// mp, win, draw, loss, point
	for k, v := range tempFileTable {
		v[3] = v[0] - v[1] - v[2]
		v[4] = v[1]*3 + v[2]
		fileTable[k] = v
	}
	return fileTable
}
