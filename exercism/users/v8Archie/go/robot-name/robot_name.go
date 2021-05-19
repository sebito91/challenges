// Package robotname solves the Exercism exercise robot_name
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const nameCount = 26 * 26 * 10 * 10 * 10
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
const debug = false

var master [nameCount]string
var register []string

type Robot struct {
	name string
}

// Name returns the Robot's name, if it has one. Otherwise it gets a new name for the robot
// and returns that.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(register) == 0 {
			return "", fmt.Errorf("Name space exhausted")
		}
		r.Reset()
	}
	return r.name, nil
}

// Reset gets a new name for the robot by picking one at random from the slice of available names.
// The chosem name is then removed from the slice. The original underlying array remains at all times.
// A better solution would keep it explicit so it is easy to reset the register.
func (r *Robot) Reset() {
	if debug {
		fmt.Printf("\rGetting integer between 0 and %6d", len(register))
	}
	i := rand.Intn(len(register))

	r.name = register[i]

	newRegister := register[:i]
	newRegister = append(newRegister, register[i+1:]...)
	register = newRegister
}

// init loads all possible names into a slice and sets the seed for random numbers
func init() {
	rand.Seed(time.Now().Unix())

	i := 0
	start := time.Now()
	for _, c1 := range alphabet {
		for _, c2 := range alphabet {
			for _, d1 := range digits {
				for _, d2 := range digits {
					for _, d3 := range digits {
						master[i] = string([]rune{c1, c2, d1, d2, d3})
						i += 1
						if debug {
							fmt.Printf("\rLoaded %6d names", i)
						}
					}
				}
			}
		}
	}

	initialiseRegister()

	if debug {
		fmt.Println()
		fmt.Println("Load completed in", time.Since(start))
	}
}

// initialiseRegister initialises the register from the master array
func initialiseRegister() {
	register = master[:]
}
