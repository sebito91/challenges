// Package igpay translates a given string into pig-latin version
package igpay

import (
	"fmt"
	"strings"
)

// testVersion is the suite we're up against
const testVersion = 1

// PigLatin takes a string and returns the pig-latin version of the input
func PigLatin(in string) (out string) {
	var a []string
	var b []string

	a = strings.Fields(in)

	for _, x := range a {
		switch {
		case strings.ContainsAny(x[:1], "aeiou") || ((x[0] == 'y' || x[0] == 'x') && !strings.ContainsAny(x[1:2], "aeiou")):
			b = append(b, fmt.Sprintf("%say", x))
		case x[:3] == "squ" || x[:3] == "thr" || x[:3] == "sch":
			b = append(b, fmt.Sprintf("%s%say", x[3:], x[:3]))
		case x[:2] == "ch" || x[:2] == "qu" || x[:2] == "th":
			b = append(b, fmt.Sprintf("%s%say", x[2:], x[:2]))
		default:
			b = append(b, fmt.Sprintf("%s%cay", x[1:], x[0]))
		}
	}

	return strings.Join(b, " ")
}
