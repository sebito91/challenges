package main

import (
	"fmt"
	"math"
)

// diff between sum of squares of first 100 natural nums minus square of sum of first 100
func diff(f int) int {
	n := float64(f)

	var other int
	for i := 1; i <= f; i++ {
		other += i * i
	}

	return int(math.Pow((n*(n+1)/2), 2)) - other
}

func main() {
	fmt.Printf("diff is: %d\n", diff(100))
}
