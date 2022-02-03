package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	output := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return output
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	output := stars + "\n" + welcomeMsg + "\n" + stars

	return output
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	temp := strings.ReplaceAll(oldMsg, "*", "")
	output := strings.TrimSpace(temp)
	return output
}
