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

var names map[string]bool

func (r *Robot) genName() {
	rand.Seed(time.Now().UTC().UnixNano())
	r.name = fmt.Sprintf("%c%c%d", charset[rand.Intn(26)], charset[rand.Intn(26)], rand.Intn(1000))
}

// Name return the new robot name
func (r *Robot) Name() string {
	if names == nil {
		names = make(map[string]bool)
	}

	if len(r.name) > 1 {
		return r.name
	}

	r.genName()
	for names[r.name] {
		r.genName()
	}
	names[r.name] = true

	return r.name
}

// Reset returns things to basics
func (r *Robot) Reset() {
	rand.Seed(time.Now().UTC().UnixNano())
	r.name = ""
}
