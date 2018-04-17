// Package gigasecond tells us when someone has lived to 10^9 seconds
package gigasecond

import (
	"math"
	"time"
)

// testVersion is our specific suite to test against
const (
	testVersion = 4
	Birthday    = "2014-12-16"
)

// AddGigasecond will take a time.Time as input, add a 10^9 second,
// and return another time.Time
func AddGigasecond(srcTime time.Time) time.Time {
	return srcTime.Add(time.Second * time.Duration(math.Pow(10, 9)))
}
