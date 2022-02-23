package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	chanL := make(chan FreqMap, len(l))
	FreqMap := FreqMap{}
	for _, v := range l {
		v := v
		go func() {
			chanL <- Frequency(v)
		}()
	}
	for range l {
		for k, v := range <-chanL {
			FreqMap[k] += v
		}
	}
	return FreqMap
}
