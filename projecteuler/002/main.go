package main

import "fmt"

// sum of even numbers in fibonacci sequence < 4M and starting with 1, 2
func main() {
	var sum int
	l, m, c := 1, 2, 2

	for {
		if c >= 4000000 {
			fmt.Printf("sum is: %d\n", sum)
			return
		}
		if c%2 == 0 {
			sum += c
		}
		m = c
		c = l + c
		l = m
	}
}
