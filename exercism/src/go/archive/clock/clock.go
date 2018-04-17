// Package clock is our add/subtract for a clock
package clock

import "fmt"

// testVersion states which specific suite we work against
const testVersion = 4

// Clock is our basic data structure representing time
type Clock struct {
	hour, min int
}

// New helps instantiate a new Clock struct
func New(hour, minute int) Clock {
	return Clock{hour: 0, min: 0}.Add((hour * 60) + minute)
}

// String is the 'stringer' method for our Clock, returning a string
// representing the time
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

// Add is a method for Clock that allows us to add/subtract
// minutes from the current time
func (c Clock) Add(minutes int) Clock {
	oneDay := 60 * 24
	totalMins := ((c.hour * 60) + c.min + minutes) % oneDay

	c.hour = (totalMins / 60) % 24
	c.min = (totalMins % 60)

	if c.hour < 0 {
		c.hour += 24
	}

	if c.min < 0 {
		if c.hour == 0 {
			c.hour = 23
		} else {
			c.hour--
		}
		c.min += 60
	}

	return c
}
