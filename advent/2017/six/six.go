package six

import "fmt"

const (
	ONE = iota
	TWO
)

// rotateKeys is a helper to redistribute keys across the list
func rotateKeys(val int) {
	keys := k1
	if val == TWO {
		keys = k2
	}

	var id int
	var m = keys[id]

	d := len(keys)

	for k, v := range keys {
		if v > m {
			m = v
			id = k
		}
	}

	keys[id] = 0

	for {
		if m == 0 {
			return
		}

		id = ((id + 1) % d)
		keys[id]++
		m--
	}
}

// DaySixPartOne of AoC 17
func DaySixPartOne() error {
	var count int
	combos := make(map[string]bool)

	for {
		if _, ok := combos[fmt.Sprint(k1)]; ok {
			fmt.Printf("DaySixPartOne codec found at: %d\n", count)
			return nil
		}

		combos[fmt.Sprint(k1)] = true
		count++
		rotateKeys(ONE)
	}
}

// DaySixPartTwo of AoC 17
func DaySixPartTwo() error {
	var count int
	combos := make(map[string]bool)

	for {
		if _, ok := combos[fmt.Sprint(k2)]; ok {
			fmt.Printf("DaySixPartOne codec found at: %d\n", count)
			return nil
		}

		combos[fmt.Sprint(k2)] = true
		count++
		rotateKeys(TWO)
	}
}

var k1 = []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}
var k2 = []int{1, 1, 0, 15, 14, 13, 12, 10, 10, 9, 8, 7, 6, 4, 3, 5}

//var keys = []int{0, 2, 7, 0}
//var keys = []int{2, 4, 1, 2}
