package clock

import (
	"fmt"
)

type Clock struct {
	hours   int
	minutes int
}

func New(hour, minute int) Clock {
	var newclock Clock
	if hour == 24 {
		newclock.hours = 0
	} else {
		newclock.hours = hour
	}
	newclock.minutes = minute
	return newclock
}

func (clk Clock) Add(minute int) Clock {
	clk.minutes += minute
	if clk.minutes >= 60 {
		clk.minutes -= 60
		clk.hours += 1
	}
	return clk
}

func (clk Clock) Subtract(minute int) Clock {
	clk.minutes -= minute
	if clk.minutes < 0 {
		clk.minutes = 0
		clk.hours -= 1
	}

	return clk
}

func (clk Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clk.hours, clk.minutes)
}
