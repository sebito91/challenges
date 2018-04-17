// Package transpose
package transpose

import "fmt"

// testVersion is the suite we're up against
const testVersion = 1

// Transpose receives
func Transpose(in []string) []string {
	fmt.Printf("in was: %+v\n", in)
	return in
}
