package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var output int
	for _, v := range birdsPerDay {
		output = output + v
	}
	return output
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var (
		output int
		k      int
	)
	k = (week - 1) * 7
	for i := k; i < k+7; i++ {
		output = output + birdsPerDay[i]
	}
	return output
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for k, v := range birdsPerDay {
		if (k+1)%2 == 1 {
			birdsPerDay[k] = v + 1
		}
	}
	return birdsPerDay
}
