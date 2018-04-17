// Package atbash implements the cipher
package atbash

import (
	"strings"
	"unicode"
)

// testVersion is the suite we're up against
const testVersion = 2

// Atbash takes in a raw string and returns the ciphered version
func Atbash(in string) (out string) {
	rot25 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'a' + (25-(r+32)+'a')%26
		case r >= 'a' && r <= 'z':
			return 'a' + (25-r+'a')%26
		case unicode.IsNumber(r):
			return r
		default:
			return ' '
		}
	}

	out = strings.Replace(strings.Map(rot25, in), " ", "", -1)

	if len(out) >= 5 {
		outer := make([]string, len(out)/5)
		var i int
		for i = 0; i+5 <= len(out); i += 5 {
			outer[i/5] = out[i : i+5]
		}

		if i < len(out) {
			outer = append(outer, out[i:])
		}

		out = strings.Join(outer, " ")
	}

	return out
}
