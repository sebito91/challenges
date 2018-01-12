package main

import (
	"fmt"
	"sort"
)

func sieve(d int) sort.IntSlice {
	var f int
	x := make(map[int]bool)

	for i := 1; i < (d + 1); i++ {
		x[i] = true
	}

	for i := 2; i < (d + 1); i++ {
		f = i * 2

		for {
			if f >= (d + 1) {
				break
			}

			x[f] = false
			f += i
		}
	}

	var out sort.IntSlice
	for a, b := range x {
		if b {
			out = append(out, a)
		}
	}
	sort.Sort(out)

	return out
}

func main() {
	x := 125000
	d := sieve(x)

	fmt.Printf("d is: %v, len: %d\n", d[10001], len(d))
}
