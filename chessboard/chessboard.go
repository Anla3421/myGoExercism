package chessboard

type (
	Rank       []bool
	Chessboard map[string]Rank
)

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var output int
	for _, v := range cb[rank] {
		if v == true {
			output = output + 1
		}
	}
	return output
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var output int
	if file < 0 || file > 8 {
		return 0
	}
	for _, rank := range cb {
		if rank[file-1] == true {
			output = output + 1
		}
	}
	return output
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var output int
	for _, rank := range cb {
		for range rank {
			output = output + 1
		}
	}
	return output
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var output int
	for i := 1; i < 9; i++ {
		count := CountInFile(cb, i)
		output = output + count
	}
	return output
}
