package main

import (
	"fmt"

	"github.com/sebito91/challenges/projecteuler/utils"
)

func main() {
	d := utils.NthPrime(1, 0, 10001)

	fmt.Printf("d is: %v, len: %d\n", d[10000], len(d))
}
