// Package hello gives us our first foray into golang exercism content
package hello

import "fmt"

// testVersion denotes which test we're runnign
const testVersion = 2

// HelloWorld is a function name to return a nice greeting to the user!
func HelloWorld(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s!", name)
	}

	return "Hello, World!"
}
