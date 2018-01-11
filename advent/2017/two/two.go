package two

import (
	"fmt"
	"sort"
)

// DayTwoPartOne of advent of code 2017
func DayTwoPartOne(in [][]int) error {
	var sum int

	for _, x := range in {
		var min = x[0]
		var max = min

		for _, y := range x {
			switch {
			case y < min:
				min = y
			case y > max:
				max = y
			}
		}
		sum += (max - min)
	}

	fmt.Printf("DayTwoPartOne checksum is: %d\n", sum)
	return nil
}

// DayTwoPartTwo of AoC 2017
func DayTwoPartTwo(in [][]int) error {
	var sum int

	for _, x := range in {
		sort.Ints(x)

		for i := 0; i < len(x); i++ {
			for j := i + 1; j < len(x); j++ {
				if x[j]%x[i] == 0 {
					sum += x[j] / x[i]
					goto Outer
				}
			}
		}
	Outer:
	}

	fmt.Printf("DayTwoPartTWo checksum is: %d\n", sum)

	return nil
}
