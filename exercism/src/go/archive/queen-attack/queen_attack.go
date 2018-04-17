// Package queenattack returns whether the queen pieces can attack one another
package queenattack

import (
	"fmt"
	"math"
)

// testVersion is the suite we're up against
const testVersion = 1

// CanQueenAttack takes the white and black positions, retuning whether an attack is possible
func CanQueenAttack(w, b string) (bool, error) {
	if len(w) == 0 || len(w) > 2 || len(b) == 0 || len(b) > 2 || w == b {
		return false, fmt.Errorf("incorrect position\n")
	}

	// error check row
	if w[0] < 'a' || w[0] > 'h' || b[0] < 'a' || b[0] > 'h' {
		return false, fmt.Errorf("row out of bounds\n")
	}

	// error check column
	if w[1] < '0' || w[1] > '7' || b[1] < '0' || b[1] > '7' {
		return false, fmt.Errorf("column out of bounds\n")
	}

	// same column or row
	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}

	// can we attack?
	if math.Abs((float64(w[0]-'a') - float64(b[0]-'a'))) == math.Abs(float64(w[1])-float64(b[1])) {
		return true, nil
	}

	// no attack possible
	return false, nil
}
