
func ConcurrentFrequency(l []string) FreqMap {
	c := make(chan FreqMap, len(l))
	for _, s := range l {
		s := s
		go func() {
			c <- Frequency(s)
		}()
	}
	counts := make(FreqMap)
	for range l {
		for k, v := range <-c {
			counts[k] += v
		}
	}
	return counts
}
