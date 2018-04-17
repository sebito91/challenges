// Package palindrome returns the list of largest palindrome products given a range
package palindrome

import (
	"fmt"
	"strconv"
)

// testVersion is the suite we're up against
const testVersion = 1

// Product is our struct containing the largest product and corresponding factorizations
type Product struct {
	Product        int
	Factorizations [][2]int
}

// checkProd is a private function to test our pair of numbers
func checkProd(a, b int) bool {
	if a*b < 0 {
		a *= -1
	}

	prod := strconv.Itoa(a * b)

	for x := 0; x < len(prod); x++ {
		if prod[x] != prod[(len(prod)-x-1)] {
			return false
		}
	}

	return true
}

// Products is our function to return the min and max palindrome given fmin/fmax range of values
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return pmin, pmax, fmt.Errorf("fmin > fmax\n")
	}

	pmin, pmax = Product{Product: fmin * fmax}, Product{}

	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			if checkProd(x, y) {
				switch {
				case x*y < pmin.Product:
					pmin.Product = x * y
					pmin.Factorizations = [][2]int{{x, y}}
				case x*y > pmax.Product:
					pmax.Product = x * y
					pmax.Factorizations = [][2]int{{x, y}}
				case x*y == pmin.Product:
					pmin.Factorizations = append(pmin.Factorizations, [2]int{x, y})
				case x*y == pmax.Product:
					pmax.Factorizations = append(pmax.Factorizations, [2]int{x, y})
				}
			}
		}
	}

	if len(pmin.Factorizations) == 0 && len(pmax.Factorizations) == 0 {
		return pmin, pmax, fmt.Errorf("No palindromes\n")
	}

	return pmin, pmax, err
}
