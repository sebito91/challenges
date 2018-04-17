// Package letter counts the letters in each given string
package letter

// FreqMap is fun too
type FreqMap map[rune]int

// Frequency function is the calculator for frequency of letters in string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency is the concurrent version of Frequency
func ConcurrentFrequency(input []string) FreqMap {
	m := FreqMap{}
	//	set := make([]FreqMap, len(input))
	set := make(chan FreqMap)

	quit := make(chan int, len(input))

	for idx, line := range input {
		go func(idx int, line string) {
			set <- Frequency(line)
			quit <- 0
		}(idx, line)
	}

	for _ = range input {
		for k, v := range <-set {
			m[k] += v
		}
	}

	for count := 0; count < len(input); count++ {
		<-quit // release the semaphore
	}

	return m
}
