package utils

import (
	"fmt"
	"sort"
)

var x = make(map[int]bool)

// NthPrime is a function that implements the Sieve of Eratosthenes
// this sieve function is unique in that the input is the number of actual
// primes to return
//
// s: the value to begin appends
// n: the number found
// e: the prime to return (eg. the xth prime in the list)
func NthPrime(s, n, e int) sort.IntSlice {
	var f int

	for i := s; i < (s * 2); i++ {
		x[i] = true
	}

	for i := 2; i < (s * 2); i++ {
		f = i * 2

		for {
			if f >= (s * 2) {
				break
			}

			x[f] = false
			f += i
		}
	}

	var out sort.IntSlice
	for a, b := range x {
		if a < s {
			continue
		}

		if b {
			out = append(out, a)
		}
	}
	sort.Sort(out)

	if len(out) < e {
		out = append(out, NthPrime((s*2)+1, len(out), e)...)
	}

	return out
}

// IsPalindrome is a helper function to checks whether a given int value
// is a palindrome
func IsPalindrome(n int) bool {
	if n < 10 && n > -10 {
		return true
	}

	d := fmt.Sprintf("%d", n)
	m := len(d)

	for i := 0; i < m/2; i++ {
		if d[i] != d[(m-1)-i] {
			return false
		}
	}
	return true
}
