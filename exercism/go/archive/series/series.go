// Package slice will handle the subnumbers within a larger number
package slice

// All returns a list of all substrings of s with length n
func All(n int, s string) (output []string) {
	if n > len(s) {
		return nil
	}

	for a := 0; a < len(s)-n+1; a++ {
		output = append(output, s[a:a+n])
	}

	return output
}

// UnsafeFirst returns the first substring of s with length n
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}

	return s[:n]
}

// First will return an ok if the string can be processed, then return the first
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}

	return s[:n], true
}
