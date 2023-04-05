// Package pythagorean returns the triplets for given numbers
package pythagorean

import "math"

// Triplet is the cast for our data
type Triplet [3]int

// Range presents the set of triplets within a range of numbers
func Range(min, max int) (output []Triplet) {
	sq := 0.0

	for outer := min; outer < max-2; outer++ {
		one := outer * outer
		for inner := outer; inner < max-1; inner++ {
			sq = math.Sqrt(float64(one + (inner * inner)))
			if math.Mod(sq, 1.0) == 0 && sq <= float64(max) {
				output = append(output, Triplet{outer, inner, int(sq)})
			}
		}
	}

	return output
}

// Sum yields the sum up to a given number p
func Sum(p int) (output []Triplet) {
	temp := Range(1, p/2)
	for _, each := range temp {
		if (each[0] + each[1] + each[2]) == p {
			output = append(output, each)
		}
	}

	return output
}