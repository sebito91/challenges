// Package diamond will print out diamond-shaped fun!
package diamond

import (
	"fmt"
	"strings"
)

const testVersion = 1

// Gen receives a char and returns a diamond up to char
func Gen(char byte) (out string, err error) {
	if !(char <= 'Z' && char >= 'A') {
		return out, fmt.Errorf("char outside of range: %c", char)
	}

	data := []string{fmt.Sprintf("%s%c%s\n", strings.Repeat(" ", int(char-'A')),
		'A', strings.Repeat(" ", int(char-'A'))),
	}

	for i := 1; i <= int(char-'A'); i++ {
		data = append(data, fmt.Sprintf("%s%c%s%c%s\n",
			strings.Repeat(" ", int(char-'A')-i), 'A'+i,
			strings.Repeat(" ", (2*i)-1), 'A'+i, strings.Repeat(" ", int(char-'A')-i)))
	}

	for i := (int(char-'A') - 1); i >= 0; i-- {
		data = append(data, data[i])
	}

	return strings.Join(data, ""), err
}
