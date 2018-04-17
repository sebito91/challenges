// Package pangram will handle our utf-8 encoded checks
package pangram

import "unicode"

// targetTestVersion is our test suite we're up against
const testVersion = 1

// IsPangram is the function to check/confirm whether we're in a pangram sentence
func IsPangram(input string) (ret bool) {
	var tester = make(map[rune]bool, 26)

	for x := 0; x < 26; x++ {
		tester[rune(x+'a')] = false
	}

	for _, each := range input {
		if unicode.IsLetter(each) {
			if _, ok := tester[unicode.ToLower(each)]; ok {
				tester[unicode.ToLower(each)] = true
			}
		}
	}

	for _, each := range tester {
		if !each {
			return false
		}
	}

	return true
}
