// Package robotname is meant to generate a unique name for a robot!
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	testVersion = 1
	charset     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Robot is an empty struct, don't need anything
type Robot struct {
	name string
}

// Name return the new robot name
func (r *Robot) Name() string {
	if len(r.name) > 1 {
		return r.name
	}

	rand.Seed(time.Now().UTC().UnixNano())
	r.name = fmt.Sprintf("%c%[1]c%d%[2]d%[2]d", charset[rand.Intn(26)], rand.Intn(10))
	return r.name
}

// Reset returns things to basics
func (r *Robot) Reset() {
	rand.Seed(time.Now().UTC().UnixNano())
	r.name = ""
}
