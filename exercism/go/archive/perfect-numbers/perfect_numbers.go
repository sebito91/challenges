// Package perfect will return whether a number is perfect, abundant or deficient
package perfect

import "fmt"

// Classification is our typedef for the perfect numbers
type Classification int

// testVersion is the suite we're up against
const (
	testVersion             = 1
	ClassificationDeficient = 0
	ClassificationAbundant  = 1
	ClassificationPerfect   = 2
)

// ErrOnlyPositive is an error message we use throughout
var ErrOnlyPositive = fmt.Errorf("Only positive numbers please!")

// Classify will return whether the number is perfect, abundant or deficient based on input
func Classify(input uint64) (name Classification, err error) {
	if input <= 0 {
		return name, ErrOnlyPositive
	}

	var sum uint64 = 1

	for x := uint64(2); x < (input/2) && (input/x) > x; x++ {
		if (input % x) == 0 {
			sum += (x + (input / x))
		}
	}

	if sum == input && input != 1 {
		return ClassificationPerfect, nil
	} else if sum > input {
		return ClassificationAbundant, nil
	}

	return ClassificationDeficient, err
}
