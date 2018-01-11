// Package letter does some stuff
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

// ConcurrentFrequency is a solid handler of ours
func ConcurrentFrequency(docs []string) FreqMap {
	ch := make(chan FreqMap)
	tot := FreqMap{}
	for _, doc := range docs {
		go func(d string) {
			ch <- Frequency(d)
		}(doc)
	}

	for range docs {
		for k, v := range <-ch {
			tot[k] += v
		}
	}

	return tot
}
