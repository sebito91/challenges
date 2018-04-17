// Package anagram detects whether the given list contains reorderings of the source material
package anagram

import (
	"sort"
	"strings"
)

// checkString is a helper function to check the inners of each string
func checkString(a, b []string) bool {
	for x := 0; x < len(a); x++ {
		if a[x] != b[x] {
			return false
		}
	}

	return true
}

// Detect is a function that receives a source + list of alternates, returning those that match
func Detect(src string, lister []string) (output []string) {
	src = strings.ToLower(src)
	data := strings.Split(src, "")
	sort.Sort(sort.StringSlice(data))

	var inner []string

	for _, each := range lister {
		each = strings.ToLower(each)
		if len(src) != len(each) {
			continue
		}

		inner = strings.Split(each, "")
		sort.Sort(sort.StringSlice(inner))

		if checkString(data, inner) && each != src {
			output = append(output, each)
		}
	}

	return output
}
