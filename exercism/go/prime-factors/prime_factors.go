// Package prime will return the list of prime factors of a given value
package prime

// testVersion is the suite we're up against
const testVersion = 2

// Factors returns a slice of prime factors (int64)
func Factors(i int64) (o []int64) {
	switch i {
	case 1:
		return []int64{}
	case 2, 3:
		return []int64{i}
	default:
		for k := int64(2); k <= i; k++ {
			if i%k != 0 {
				continue
			}

			for i%k == 0 {
				o = append(o, k)
				i = i / k
			}
		}
	}

	return o
}
