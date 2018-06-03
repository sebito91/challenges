// Package twofer prints out the saying
package twofer

import "fmt"

// ShareWith receives a name value and outputs a string with that name
func ShareWith(in string) string {
	if len(in) < 1 {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", in)
}
