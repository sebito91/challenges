// Package pascal returns the tree of values for Pascal's fun triangle
package pascal

// Triangle does the needful
func Triangle(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	if n == 2 {
		return [][]int{{1}, {1, 1}}
	}

	// if n > 1, recursion!
	prev := Triangle(n - 1)
	curr := make([]int, n)
	curr[0], curr[n-1] = 1, 1

	for count := 1; count <= (n / 2); count++ {
		curr[count] = prev[n-2][count-1] + prev[n-2][count]
		curr[n-count-1] = curr[count]
	}

	return append(prev, curr)
}
