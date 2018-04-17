// Package stringset applies a series of functions to a given set of items
package stringset

import (
	"fmt"
	"sort"
	"strings"
)

// testVersion is the version of the test-suite we're up against
const testVersion = 4

// Set is the default type for our function
type Set struct {
	m map[string]bool
}

// New returns an empty instance of the Set type
func New() Set {
	return Set{m: make(map[string]bool)}
}

// String is the string repr of a Set
func (s Set) String() string {
	var d []string

	for k := range s.m {
		d = append(d, fmt.Sprintf("\"%s\"", k))
	}

	return fmt.Sprintf("{%s}", strings.Join(d, ", "))
}

// IsEmpty tests whether our given set is empty
func (s Set) IsEmpty() bool {
	return len(s.m) == 0
}

// NewFromSlice returns a new set from the provided string slice
func NewFromSlice(ins []string) Set {
	s := Set{m: make(map[string]bool)}

	for _, k := range ins {
		if _, ok := s.m[k]; !ok {
			s.m[k] = true
		}
	}

	return s
}

// Has checks whether the given set contains a specific string
func (s Set) Has(in string) bool {
	_, ok := s.m[in]
	return ok
}

// Subset checks whether two given slices are a subset of one another
func Subset(s1, s2 Set) bool {
	for k := range s1.m {
		if _, ok := s2.m[k]; !ok {
			return false
		}
	}

	return true
}

func (s Set) getkeys() []string {
	var out []string

	for k := range s.m {
		out = append(out, k)
	}
	sort.Strings(out)

	return out
}

// Disjoint checks whether two gives Sets are disjoint, i.e. without common elements
func Disjoint(s1, s2 Set) bool {
	a := s1.getkeys()
	b := s2.getkeys()

	for _, k := range a {
		if _, ok := s2.m[k]; ok {
			return false
		}
	}

	for _, k := range b {
		if _, ok := s1.m[k]; ok {
			return false
		}
	}

	return true
}

// Equal checks whether two given Sets are equal
func Equal(s1, s2 Set) bool {
	if len(s1.m) != len(s2.m) {
		return false
	}

	a := s1.getkeys()
	b := s2.getkeys()

	// wanted to avoid using reflect
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Add will add a new string to the Set
func (s Set) Add(k string) {
	if _, ok := s.m[k]; !ok {
		s.m[k] = true
	}
}

// Intersection returns the common elements from the given Sets
func Intersection(s1, s2 Set) Set {
	i := make(map[string]bool)

	for _, k := range s1.getkeys() {
		if _, ok := s2.m[k]; ok {
			i[k] = true
		}
	}

	return Set{m: i}
}

// Difference returns a Set of non-common elements from both given Sets
func Difference(s1, s2 Set) Set {
	d := make(map[string]bool)

	for _, k := range s1.getkeys() {
		if _, ok := s2.m[k]; !ok {
			d[k] = true
		}
	}

	return Set{m: d}
}

// Union returns the combined output from both given sets
func Union(s1, s2 Set) Set {
	u := make(map[string]bool)

	for _, c := range [][]string{s1.getkeys(), s2.getkeys()} {
		for _, k := range c {
			if _, ok := u[k]; !ok {
				u[k] = true
			}
		}
	}

	return Set{m: u}
}
