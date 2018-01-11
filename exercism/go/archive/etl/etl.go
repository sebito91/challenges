// Package etl gives us the letter breakdown for each element
package etl

import "strings"

// Transform inverts the letters per score and returns scores per letter
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)

	for k, v := range input {
		for _, j := range v {
			output[strings.ToLower(j)] = k
		}
	}

	return output
}
