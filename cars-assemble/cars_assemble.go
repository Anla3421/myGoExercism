package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	output := float64(productionRate) * successRate * 0.01
	return output
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	output := int(float64(productionRate) * successRate * 0.01 / 60)
	return output
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	quotient := carsCount / 10
	remainder := carsCount % 10
	output := uint(quotient*95000 + remainder*10000)
	return output
}
