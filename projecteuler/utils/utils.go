package utils

import "sort"

var x = make(map[int]bool)

// sieve is a function that implements the Sieve of Eratosthenes
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
