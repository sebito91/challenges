// Package foodchain spits out the weirdest song lyrics ever
package foodchain

import "strings"

// testVersion is the test suite we're up against
const (
	testVersion = 2
	verseCloser = "I don't know why she swallowed the fly. Perhaps she'll die."
	songCloser  = "She's dead, of course!"
)

// Verse is the function to return the list of verses for the song
func Verse(verse int) string {
	var output []string

	openers := map[int]string{
		1: "fly",
		2: "spider",
		3: "bird",
		4: "cat",
		5: "dog",
		6: "goat",
		7: "cow",
		8: "horse",
	}

	quals := map[int]string{
		1: verseCloser,
		2: "It wriggled and jiggled and tickled inside her.",
		3: "How absurd to swallow a bird!",
		4: "Imagine that, to swallow a cat!",
		5: "What a hog, to swallow a dog!",
		6: "Just opened her throat and swallowed a goat!",
		7: "I don't know how she swallowed a cow!",
	}

	output = append(output, ("I know an old lady who swallowed a " + openers[verse] + "."))

	if verse > 1 && verse != 8 {
		output = append(output, quals[verse])

		for count := verse; count > 1; count-- {
			str := "She swallowed the " + openers[count] + " to catch the " + openers[count-1]
			if count == 3 {
				str += strings.Replace(quals[count-1], "It", " that", 1)
			} else {
				str += "."
			}
			output = append(output, str)
		}
	}

	if verse != 8 {
		output = append(output, verseCloser)
	} else {
		output = append(output, songCloser)
	}

	return strings.Join(output, "\n")
}

// Verses is the combo for verse
func Verses(a, b int) string {
	return Verse(a) + "\n\n" + Verse(b)
}

// Song is the collection of all terrible verses
func Song() string {
	var output []string
	for each := 1; each < 9; each++ {
		output = append(output, Verse(each))
	}

	return strings.Join(output, "\n\n")
}
