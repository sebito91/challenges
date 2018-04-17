// Package grains returns the total grains per square of a chessboard
package grains

import (
	"fmt"
	"math"
	"runtime"
	"sync/atomic"
	"time"
)

// testVersion is the suite we're up against
const testVersion = 1

// Total returns the number of grains up to and including the given square
func Total() (output uint64) {
	for each := 1; each < 65; each++ {
		go func(each int) {
			if val, err := Square(each); err == nil {
				atomic.AddUint64(&output, val)
			} else {
				fmt.Printf("error: %v\n", err)
			}
			runtime.Gosched()
		}(each)
	}

	time.Sleep(time.Second)

	return atomic.LoadUint64(&output)
}

// Square returns the number of grains on the specified chessboard square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("number out of range: %+v\n", n)
	}

	return uint64(math.Pow(2.0, float64(n-1))), nil
}
