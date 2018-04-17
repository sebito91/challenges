// Package strand will return an RNA complement for a given DNA strand
package strand

import "bytes"

// testVersion is the suite we're up against
const testVersion = 3

var nukes = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

// ToRNA is a function to return an RNA complement for a given DNA strand
func ToRNA(input string) (output string) {
	var buf bytes.Buffer

	for _, each := range input {
		buf.WriteRune(nukes[each])
	}

	return buf.String()
}
