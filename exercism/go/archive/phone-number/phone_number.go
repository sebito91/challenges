// Package phonenumber cleans up input to match NANP definition
package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

// testVersion is the suite we're up against
const testVersion = 2

// Number takes in a string and returns the raw number
func Number(in string) (out string, err error) {
	f := func(r rune) bool {
		return !unicode.IsNumber(r)
	}

	out = strings.Join(strings.FieldsFunc(in, f), "")
	switch {
	case len(out) == 11 && out[0] == '1' && out[1] >= '2' && out[4] >= '2':
		return out[1:], nil
	case len(out) == 11 && out[0] > '1':
		return "", fmt.Errorf("invalid number: %s", in)
	case len(out) < 10 || out[0] <= '1' || out[3] <= '1' || len(out) > 11:
		return "", fmt.Errorf("invalid number: %s", in)
	}

	return out, nil
}

// AreaCode takes in a number and returns just the area code portion
func AreaCode(in string) (out string, err error) {
	out, err = Number(in)
	if err != nil {
		return "", err
	}

	return out[:3], nil
}

// Format takes in a raw value and returns the properly formatted version
func Format(in string) (out string, err error) {
	out, err = Number(in)
	if err != nil {
		return "", err
	}

	out = fmt.Sprintf("(%s) %s-%s", out[:3], out[3:6], out[6:])
	return out[:], err
}
