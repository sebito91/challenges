package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./input")
	check(err)

	fmt.Printf("testing dat: %+v\n", dat)

}
