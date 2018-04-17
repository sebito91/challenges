// Package bob is the fun stuff where we reply based on a given src string
package bob // package name must match the package name in bob_test.go

import (
	"strings"
	"unicode"
)

// testVersion is the specific test suite we're confirmed working for
const testVersion = 2 // same as targetTestVersion

func checkAllcaps(theval string) bool {
	retval := false

	for _, item := range theval {
		if unicode.IsUpper(item) {
			retval = true
			continue
		}

		if unicode.IsPunct(item) || unicode.IsSpace(item) || unicode.IsNumber(item) || unicode.IsSymbol(item) {
			continue
		} else {
			return false
		}
	}

	return retval
}

func checkDigits(theval string) bool {
	for _, item := range theval {
		if unicode.IsPunct(item) || unicode.IsSpace(item) {
			continue
		}

		if !('0' <= item && item <= '9') {
			return false
		}
	}

	return true
}

// Hey is our witty reply to a given src string
func Hey(src string) (wittyReply string) {
	src = strings.TrimSpace(src) // get rid of spaces before we start!

	if len(src) == 0 {
		return "Fine. Be that way!"
	}

	if checkAllcaps(src) {
		return "Whoa, chill out!"
	}

	if checkDigits(src) {
		if src[len(src)-1] != '?' {
			return "Whatever."
		}

		return "Sure."
	}

	switch src[len(src)-1] {
	case '!':
		return "Whatever."
	case '?':
		return "Sure."
	}

	return "Whatever." // default reply
}
