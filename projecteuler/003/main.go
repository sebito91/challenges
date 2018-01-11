package main

import (
	"fmt"
	"math"
	"sort"
)

func sieve(d int) sort.IntSlice {
	if d <= 0 {
		return []int{}
	}

	var f int
	m := math.Sqrt(float64(d))
	d = int(m)
	x := make(map[int]bool)

	fmt.Printf("d is: %d\n", d)

	for i := 0; i < (d + 1); i++ {
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

	return out
}

// print largest prime factor for 600851475143
func main() {
	val := 600851475143
	a := sieve(val)
	sort.Sort(sort.Reverse(a))

	for _, x := range a {
		if val%x == 0 {
			fmt.Printf("largest is: %d\n", x)
			return
		}
	}
}
