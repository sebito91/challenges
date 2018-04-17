// Package triangle is a setup to return whether we have a triangle
package triangle

import "math"

// Kind is our typedef to check triangle definition
type Kind int

// testVersion is the suite we're up against
const (
	testVersion      = 3
	NaT         Kind = iota // not a triangle
	Equ                     // equilateral
	Iso                     // isosceles
	Sca                     // scalene
)

// KindFromSides will tell us whether we have a triangle or something else
func KindFromSides(a, b, c float64) Kind {
	if ((a + b + c) <= 0) || a < 0 || b < 0 || c < 0 { // initial check for NaT
		return NaT
	}

	if (math.IsInf(a, 0)) || (math.IsInf(b, 0)) || (math.IsInf(c, 0)) {
		return NaT
	}

	if (a == b) && (b == c) { // equilateral
		return Equ
	}

	// isosceles and scalene
	if ((a + b) >= c) && ((b + c) >= a) && ((a + c) >= b) {
		if (a == b) || (a == c) || (b == c) {
			return Iso
		}

		return Sca
	}

	return NaT
}
