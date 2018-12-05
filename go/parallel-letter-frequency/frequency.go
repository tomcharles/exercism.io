package letter

// FreqMap represents the frequency of letters in a string
type FreqMap map[rune]int

// Frequency roll a string up into a FreqMap
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency roll a list of strings up into a FreqMap concurrently
func ConcurrentFrequency(words []string) FreqMap {
	c := make(chan FreqMap)
	for _, w := range words {
		go func(word string) { c <- Frequency(word) }(w)
	}

	fm := FreqMap{}
	for range words {
		for k, v := range <-c {
			fm[k] += v
		}
	}
	return fm
}
