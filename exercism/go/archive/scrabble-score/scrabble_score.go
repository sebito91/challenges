// Package scrabble gives us the score for a given scrabble string
package scrabble

import "strings"

// testVersion is the suite we're up against
const testVersion = 4

// Score returns the scrabble score for a given string
func Score(input string) (score int) {
	for _, each := range strings.ToLower(input) {
		switch each {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		default:
			continue
		}
	}

	return score
}
