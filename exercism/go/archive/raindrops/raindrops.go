// Package raindrops is a factoring function that tells us if we're divisible by 3, 5, 7 inclusive or none at all!
package raindrops

import (
	"bytes"
	"strconv"
)

// testVersion confirms the version we're testing against
const testVersion = 2

// Convert is our function to handle the input (int), and convert to a string based on whether its divisible
// by 3, 5, 7 inclusive (or none at all)
func Convert(theval int) string {
	var output bytes.Buffer

	if (theval % 3) == 0 {
		output.WriteString("Pling")
	}

	if (theval % 5) == 0 {
		output.WriteString("Plang")
	}

	if (theval % 7) == 0 {
		output.WriteString("Plong")
	}

	if output.Len() == 0 {
		return strconv.Itoa(theval)
	}

	return output.String()
}
