// Package romannumerals returns the roman-numeral equiv of a given value
package romannumerals

import (
	"fmt"
	"strings"
)

// testVersion is the suite we're up against
const testVersion = 3

// addChars is a private function to handle the specific chars in a section
func addChars(div int, low, mid, high string) string {
	var data []string

	switch {
	case div > 0 && div < 4:
		data = append(data, strings.Repeat(low, div))
	case div == 4:
		data = append(data, fmt.Sprintf("%s%s", low, mid))
	case div == 5:
		data = append(data, mid)
	case div > 5 && div <= 8:
		data = append(data, mid)
		data = append(data, strings.Repeat(low, div-5))
	case div == 9:
		data = append(data, fmt.Sprintf("%s%s", low, high))
	}

	return strings.Join(data, "")
}

// ToRomanNumeral will convert the given arabic int to a roman_numeral string
func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", fmt.Errorf("number is outside of range: %d\n", input)
	}

	var data []string
	data = append(data, strings.Repeat("M", input/1000))
	input %= 1000

	data = append(data, addChars(input/100, "C", "D", "M"))
	input %= 100

	data = append(data, addChars(input/10, "X", "L", "C"))
	input %= 10

	data = append(data, addChars(input, "I", "V", "X"))

	return strings.Join(data, ""), nil
}
