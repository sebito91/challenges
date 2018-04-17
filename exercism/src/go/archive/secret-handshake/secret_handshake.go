// Package secret is the handshake generator for a given number
package secret

func reverse(input []string) []string {
	for count := len(input)/2 - 1; count >= 0; count-- {
		rev := len(input) - 1 - count
		input[count], input[rev] = input[rev], input[count]
	}

	return input
}

// Handshake translates a number into a series of secret moves
func Handshake(number int) (output []string) {
	if number <= 0 || number == 32 {
		return output
	}

	//	tester := number % 32
	notes := map[int]string{
		1: "wink",
		2: "double blink",
		3: "close your eyes",
		4: "jump",
	}

	for count := 0; count < 4; count++ {
		if ((number >> uint(count)) & 1) == 1 {
			output = append(output, notes[count+1])
		}
	}

	if number >= 16 {
		return reverse(output)
	}

	return output
}
