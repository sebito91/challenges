package main

import (
	"fmt"

	"github.com/sebito91/challenges/projecteuler/utils"
)

func main() {
	var max int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if ok := utils.IsPalindrome(i * j); ok && (i*j > max) {
				max = i * j
			}
		}
	}

	fmt.Printf("max 3-digit palindrome is: %d\n", max)
}
