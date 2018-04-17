package reverse

import (
	"strings"
)

// String reverse a given input string value
func String(in string) string {
	if len(in) == 0 {
		return in
	}

	d := strings.Split(in, "")

	w := len(d) - 1
	var c string
	for x := 0; x < w/2; x++ {
		c = d[x]
		d[x] = d[w-x]
		d[w-x] = c
	}

	return strings.Join(d, "")
}
