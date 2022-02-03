package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime int = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven == 0 {
		panic("RemainingOvenTime not implemented")
	}
	remain := OvenTime - actualMinutesInOven
	return remain
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	if numberOfLayers == 0 {
		panic("PreparationTime not implemented")
	}
	prepare := numberOfLayers * 2
	return prepare
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	if actualMinutesInOven == 0 {
		panic("ElapsedTime not implemented")
	}
	elapsed := PreparationTime(numberOfLayers) + actualMinutesInOven
	return elapsed
}
