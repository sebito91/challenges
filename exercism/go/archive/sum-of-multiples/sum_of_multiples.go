// Package summultiples will take a value and return the sum of the multiples
package summultiples

// Func cast our function defintion for the sneaky MultipleSummer return
type Func func(int) int

// MultipleSummer is a variadic function because exercism is mean
func MultipleSummer(nums ...int) Func {
	test := func(n int) (score int) {
		keys := make(map[int]bool)

		for _, each := range nums {
			for count := 1; (each * count) < n; count++ {
				if _, ok := keys[each*count]; ok {
					continue
				}
				keys[each*count] = true
				score += (each * count)
			}
		}

		return score
	}

	return test
}
