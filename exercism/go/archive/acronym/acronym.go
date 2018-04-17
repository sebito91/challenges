// Package acronym will convert the given string into short-form
package acronym

import (
	"strings"
	"unicode"
)

// testVersion is the test suite we're up against
const testVersion = 1

// abbreviate takes in a string, formats its short-form and returns it
func abbreviate(input string) string {
	var output []byte

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	if strings.Contains(input, ":") {
		return strings.Split(input, ":")[0]
	}

	thelist := strings.FieldsFunc(input, f)

	for _, each := range thelist {
		output = append(output, strings.ToUpper(each)[0])

		for _, item := range each[1:] {
			if unicode.IsUpper(item) {
				output = append(output, byte(item))
			}
		}
	}

	return string(output)
}
