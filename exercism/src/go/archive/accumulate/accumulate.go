// Package accumulate is a nice module to apply a specific function to each provided element
package accumulate

// testVersion is the suite we're up against
const testVersion = 1

// Accumulate applies a given function to the provided inputs, returning their answers
func Accumulate(inputs []string, theFunc func(string) string) (output []string) {
	for _, each := range inputs {
		output = append(output, theFunc(each))
	}

	return output
}
