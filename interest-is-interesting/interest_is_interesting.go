package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var output float32
	switch {
	case balance < 0:
		output = 3.213
	case balance >= 0 && balance < 1000:
		output = 0.5
	case balance >= 1000 && balance < 5000:
		output = 1.621
	case balance >= 5000:
		output = 2.475
	}
	return output
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	output := balance * float64(InterestRate(balance)) * 0.01
	return output
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	output := balance * (1 + float64(InterestRate(balance))*0.01)
	return output
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	rate := InterestRate(balance)
	for years = 1; ; years++ {
		output := balance * (1 + float64(rate)*0.01)
		if output >= targetBalance {
			break
		}
		rate = InterestRate(output)
		balance = output
	}
	return years
}
