// Package luhn checks whether a given number is a valid Luhn formula number
package luhn

import (
	"bytes"
	"strconv"
	"unicode"
)

// AddCheck function will add a digit to make the provided input a valid Luhn number
func AddCheck(input string) string {
	if len(input) <= 0 {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString(input)

	for x := 0; x < 10; x++ {
		buf.WriteString(strconv.Itoa(x))

		if Valid(buf.String()) {
			return buf.String()
		}

		buf.Reset()
		buf.WriteString(input)
	}

	return ""
}

// Valid function will determine whether the provided input is a valid Luhn
func Valid(input string) bool {
	if len(input) <= 0 {
		return false
	}

	last, sum := false, 0
	for x := (len(input) - 1); x >= 0; x-- {
		if !unicode.IsNumber(rune(input[x])) {
			continue
		}

		if last {
			if val := int(input[x]-'0') * 2; val >= 10 {
				sum += (val - 9)
			} else {
				sum += val
			}
			last = false
		} else {
			sum += int(input[x] - '0')
			last = true
		}
	}

	if (sum%10) != 0 || sum == 0 {
		return false
	}

	return true
}
