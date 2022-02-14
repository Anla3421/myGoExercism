package raindrops

import "strconv"

func Convert(number int) string {
	switch {
	case number%(3*5*7) == 0:
		return "PlingPlangPlong"
	case number%(5*7) == 0:
		return "PlangPlong"
	case number%(3*7) == 0:
		return "PlingPlong"
	case number%(3*5) == 0:
		return "PlingPlang"
	case number%3 == 0:
		return "Pling"
	case number%5 == 0:
		return "Plang"
	case number%7 == 0:
		return "Plong"
	}
	return strconv.Itoa(number)
}
