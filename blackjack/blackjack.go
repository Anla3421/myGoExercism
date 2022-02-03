package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	table := map[string]int{"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7,
		"eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10, "other": 0}
	return table[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	if ParseCard(card1)+ParseCard(card2) != 21 {
		return false
	}
	return true
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	// - Stand (S)
	// - Hit (H)
	// - Split (P)
	// - Automatically win (W)
	switch {
	case isBlackjack == false:
		return "P"
	case isBlackjack == true && (dealerScore == 11 || dealerScore == 10):
		return "S"
	case isBlackjack == true && dealerScore != 11:
		return "W"
	default:
		return "S"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case handScore >= 12 && handScore <= 16:
		if dealerScore >= 7 {
			return "H"
		}
		return "S"
	default:
		return "S"
	}
}
