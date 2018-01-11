// Package strain defines a few functions on collections
package strain

// testVersion is the suite we're up against
const testVersion = 1

// Ints is a collection of int values
type Ints []int

// Lists is a collection of Ints
type Lists []Ints

// Strings is a collection of string values
type Strings []string

// Keep is a bound method that processes a keep function on a set of Ints
func (i Ints) Keep(f func(int) bool) (r Ints) {
	for _, x := range i {
		if f(x) {
			r = append(r, x)
		}
	}
	return r
}

// Discard is a bound method that processes a discard function on a set of Ints
func (i Ints) Discard(f func(int) bool) (r Ints) {
	for _, x := range i {
		if !f(x) {
			r = append(r, x)
		}
	}
	return r
}

// Keep is a bound method that processes a keep function on a list of Ints
func (l Lists) Keep(f func([]int) bool) (r Lists) {
	for _, x := range l {
		if f(x) {
			r = append(r, x)
		}
	}
	return r
}

// Keep is a bound method that processes a keep function on a list of strings
func (s Strings) Keep(f func(string) bool) (r Strings) {
	for _, x := range s {
		if f(x) {
			r = append(r, x)
		}
	}
	return r
}
