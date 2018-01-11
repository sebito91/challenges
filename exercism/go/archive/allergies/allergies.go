// Package allergies returns the list of ailments from a given score
package allergies

import (
	"fmt"
	"strconv"
)

// testVersion is the suite we're up against
const testVersion = 1

var items = [8]string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns the list of ailments given an allergy "score"
func Allergies(input uint) (ailments []string) {
	score := fmt.Sprintf("%08s", strconv.FormatUint(uint64(input%256), 2))

	for x := len(score) - 1; x >= 0; x-- {
		if score[x] == '1' {
			ailments = append(ailments, items[len(score)-x-1])
		}
	}

	return ailments
}

// AllergicTo returns whether the given allergy "score" matches the ailment
func AllergicTo(input uint, ailment string) bool {
	for _, each := range Allergies(input) {
		if each == ailment {
			return true
		}
	}

	return false
}
