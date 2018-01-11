// Package hamming is the packing to calculate the 'hamming distance' in a DNA strand
package hamming

import "fmt"

// testVersion is our confirmation that we're running the right test series
const testVersion = 4

// Distance is the function that'll return the number of errors in our DNA strand!
func Distance(a, b string) (int, error) {
	failcount := 0
	if len(a) != len(b) {
		return -1, fmt.Errorf("strings are not the same length\n")
	}

	for count := 0; count < len(a); count++ {
		if a[count] != b[count] {
			failcount++
		}
	}

	return failcount, nil
}
