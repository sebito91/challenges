// Package leap helps us calculate whether the input is a leap year
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear is great for telling us whether the int year is a leap (return bool)
func IsLeapYear(year int) bool {
	if year == 0 {
		return false
	}

	if ((year % 4) == 0) && (((year % 100) != 0) || ((year % 400) == 0)) {
		return true
	}

	return false
}
