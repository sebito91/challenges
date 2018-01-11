// Package say writes out the text for a given value input
package say

import (
	"fmt"
	"math"
)

var nums = map[uint64]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

// Say is a function that returns the string format of a given input uint64
func Say(input uint64) string {
	if _, ok := nums[input]; ok {
		return fmt.Sprintf("%s", nums[input])
	}

	switch {
	case input < 100: // tens
		return fmt.Sprintf("%s-%s", nums[(input/10)*10], nums[input%10])
	case input < 1000: // hundreds
		if (input % 100) == 0 {
			return fmt.Sprintf("%s hundred", nums[input/100])
		}

		return fmt.Sprintf("%s hundred %s", nums[input/100], Say(input%100))
	case input < 1e6: // thousands
		if (input % 1000) == 0 {
			return fmt.Sprintf("%s thousand", Say(input/1000))
		}

		return fmt.Sprintf("%s thousand %s", Say(input/1000), Say(input%1000))
	case input < 1e9: // millions
		if (input % 1e6) == 0 {
			return fmt.Sprintf("%s million", Say(input/1e6))
		}

		return fmt.Sprintf("%s million %s", Say(input/1e6), Say(input%1e6))
	case input < 1e12: // billions
		if (input % 1e9) == 0 {
			return fmt.Sprintf("%s billion", Say(input/1e9))
		}

		return fmt.Sprintf("%s billion %s", Say(input/1e9), Say(input%1e9))
	case input < 1e15: // trillions
		if (input % 1e12) == 0 {
			return fmt.Sprintf("%s trillion", Say(input/1e12))
		}

		return fmt.Sprintf("%s trillion %s", Say(input/1e12), Say(input%1e12))
	case input < 1e18: // quadrillions
		if (input % 1e15) == 0 {
			return fmt.Sprintf("%s quadrillion", Say(input/1e15))
		}

		return fmt.Sprintf("%s quadrillion %s", Say(input/1e15), Say(input%1e15))
	case input <= math.MaxUint64:
		if (input % 1e18) == 0 {
			return fmt.Sprintf("%s quintillion", Say(input/1e18))
		}

		return fmt.Sprintf("%s quintillion %s", Say(input/1e18), Say(input%1e18))
	}

	return ""
}
