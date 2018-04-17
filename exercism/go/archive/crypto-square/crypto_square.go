// Package cryptosquare returns the square-coded version of a string
package cryptosquare

import (
	"bytes"
	"math"
	"strings"
	"unicode"
)

// testVersion is the test suite we're up against
const testVersion = 2

func round(f float64) float64 {
	return math.Floor(f + .5)
}

// Encode receives a decoded string and returns the encoded version using square-coding
func Encode(input string) (output string) {
	if len(input) <= 0 {
		return ""
	}

	var buf bytes.Buffer

	for _, each := range strings.ToLower(input) {
		if !unicode.IsLetter(each) && !unicode.IsNumber(each) {
			continue
		}
		buf.WriteRune(each)
	}

	input = buf.String()
	buf.Reset()

	rows := int(round(math.Sqrt(float64(len(input)))))
	cols := int(math.Ceil(float64(len(input) / rows)))

	if (len(input) % rows) != 0 {
		cols++
	}

	slicer := []string{}

	for x := 0; x < cols; x++ {
		for y := 0; y < rows && (((cols * y) + x) < (len(input))); y++ {
			buf.WriteByte(input[(cols*y)+x])
		}
		slicer = append(slicer, buf.String())
		buf.Reset()
	}

	return strings.Join(slicer, " ")
}
