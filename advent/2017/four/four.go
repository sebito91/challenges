package four

import (
	"sort"
)

type sorter []rune

func (s sorter) Len() int           { return len(s) }
func (s sorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sorter) Less(i, j int) bool { return s[i] < s[j] }

// DayFourPartOne AoC '17
func DayFourPartOne(k [][]string) int {
	var total int
	for _, x := range k {
		if len(x) == 0 {
			continue
		}

		m := make(map[string]bool)

		for _, y := range x {
			if _, ok := m[y]; ok {
				goto start
			}
			m[y] = true
		}
		total++
	start:
	}

	return total
}

// sortString is a bit of slice magic :D
func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sorter(r))
	return string(r)
}

// DayFourPartTwo AoC '17
func DayFourPartTwo(k [][]string) int {
	var d = make([][]string, len(k))

	// sanitize inputs
	for _, x := range k {
		var f []string
		for _, s := range x {
			f = append(f, sortString(s))
		}
		sort.Strings(f)
		d = append(d, f)
	}

	return DayFourPartOne(d)
}
