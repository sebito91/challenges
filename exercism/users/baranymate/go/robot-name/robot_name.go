// Package robotname implements the solution to the exercism problem Robot Name
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	ID string
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Name returns the name of the robot
func (r *Robot) Name() (string, error) {
	if r.ID == "" {

		r.ID = r.generateName()
	}

	return r.ID, nil
}

// Reset resets the name of the robot
func (r *Robot) Reset() {
	r.ID = r.generateName()
}

// generateName generates a name for the robot
func (r Robot) generateName() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	runes := make([]rune, 2)
	runes[0] = letters[rand.Intn(len(letters))]
	runes[1] = letters[rand.Intn(len(letters))]

	temp := string(runes)

	return fmt.Sprintf("%s%03d", temp, random.Intn(999))
}
