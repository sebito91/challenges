package three

import (
	"fmt"
	"math"
)

const (
	up int = iota
	down
	left
	right
)

// DayThreePartOne of AoC 2017
func DayThreePartOne(in int) error {
	sq := int(math.Ceil(math.Sqrt(float64(in))) + 1)
	if sq < 10 {
		sq = 10
	}

	var table = make([][]int, sq)

	for i := 0; i < sq; i++ {
		table[i] = make([]int, sq)
	}

	x := sq / 2
	y := sq / 2
	step := 1
	table[y][x] = 1 // center

	dir := right

	for {
		if step == in {
			break
		}
		step++

		switch dir {
		case right:
			x++
			table[y][x] = step

			if table[y-1][x] == 0 {
				dir = up
			}

		case left:
			x--
			table[y][x] = step

			if table[y+1][x] == 0 {
				dir = down
			}

		case up:
			y--
			table[y][x] = step

			if table[y][x-1] == 0 {
				dir = left
			}

		case down:
			y++
			table[y][x] = step

			if table[y][x+1] == 0 {
				dir = right
			}
		}
		//		printTable(table)
		//		fmt.Printf("y: %d, x: %d, dir: %d, step: %d\n", y, x, dir, step)
	}

	diff := (math.Abs(float64((x - (sq / 2)))) + math.Abs(float64((y - (sq / 2)))))
	fmt.Printf("DayThreePartOne checksum is: %d, %0.0f\n", in, diff)
	return nil
}

// printTable is here for debugging
// NOTE: this table is inverted meaning (0,0) is the top-left corner
// instead of bottom-right! All coordinates are based from the top-left
//
// To check this out for small numbers (~10000) enable the printTable functions
//func printTable(t [][]int) {
//	for _, y := range t {
//		fmt.Printf("[ ")
//		for _, v := range y {
//			fmt.Printf("%5d ", v)
//		}
//		fmt.Printf("]\n")
//	}
//}

// DayThreePartTwo of AoC 2017
func DayThreePartTwo(in int) error {
	sq := int(math.Ceil(math.Sqrt(float64(in))))
	if sq < 10 {
		sq = 10
	}
	var table = make([][]int, sq)
	var sum int

	for i := 0; i < sq; i++ {
		table[i] = make([]int, sq)
	}

	x := sq / 2
	y := sq / 2
	step := 1
	table[y][x] = 1 // center
	//	printTable(table)

	dir := right

	for {
		if sum > in {
			break
		}

		switch dir {
		case right:
			// sum += current + one-up + one-up & one-right + one-up & two-right
			sum = table[y][x] + table[y-1][x] + table[y-1][x+1] + table[y-1][x+2]
			x++
			table[y][x] = sum

			if table[y-1][x] == 0 {
				dir = up
			}

		case left:
			// sum += current + one-down + one-left & one-down + two-left & one-down
			sum = table[y][x] + table[y+1][x] + table[y+1][x-1] + table[y+1][x-2]
			x--
			table[y][x] = sum

			if table[y+1][x] == 0 {
				dir = down
			}

		case up:
			// sum += current + one-left + one-left & one-up + one-left & two-up
			sum = table[y][x] + table[y][x-1] + table[y-1][x-1] + table[y-2][x-1]
			y--
			table[y][x] = sum

			if table[y][x-1] == 0 {
				dir = left
			}

		case down:
			// sum += current + one-right + one-right & one-down + one-right & two-down
			sum = table[y][x] + table[y][x+1] + table[y+1][x+1] + table[y+2][x+1]
			y++
			table[y][x] = sum

			if table[y][x+1] == 0 {
				dir = right
			}
		}
		step++
		//		printTable(table)
		//		fmt.Printf("sum is: %d, y: %d, x: %d, dir: %d, step: %d\n", sum, y, x, dir, step)
	}

	fmt.Printf("DayThreePartTwo checksum is: %d, %d\n", in, sum)

	return nil
}
