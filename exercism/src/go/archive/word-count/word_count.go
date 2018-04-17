// Package wordcount returns the number of occurrences of each word in a phrase
package wordcount

import (
	"strings"
	"unicode"
)

// testVersion is the suite we're up against
const testVersion = 3

// Frequency is a typedef for our map
type Frequency map[string]int

// WordCount returns the frequency of each word in the given phrase
func WordCount(phrase string) (freq Frequency) {
	freq = make(Frequency)
	data := strings.FieldsFunc(strings.ToLower(phrase), func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	})

	for _, each := range data {
		freq[strings.TrimFunc(each, func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})]++
	}

	return freq
}
