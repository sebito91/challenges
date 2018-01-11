// Package twelve is our very holiday-themed package to return a song
package twelve

import "bytes"

// testVersion represents the test suite we're up against
const testVersion = 1

/// songDay is the basic structure of each day in the song
type songDay struct {
	day  string
	gift string
}

// global var for us to use
var song = make(map[int]songDay, 12)

// initSong is an internal function to initialize our song
func initSong() bool {
	song[0] = songDay{day: "first", gift: "a Partridge in a Pear Tree."}
	song[1] = songDay{day: "second", gift: "two Turtle Doves"}
	song[2] = songDay{day: "third", gift: "three French Hens"}
	song[3] = songDay{day: "fourth", gift: "four Calling Birds"}
	song[4] = songDay{day: "fifth", gift: "five Gold Rings"}
	song[5] = songDay{day: "sixth", gift: "six Geese-a-Laying"}
	song[6] = songDay{day: "seventh", gift: "seven Swans-a-Swimming"}
	song[7] = songDay{day: "eighth", gift: "eight Maids-a-Milking"}
	song[8] = songDay{day: "ninth", gift: "nine Ladies Dancing"}
	song[9] = songDay{day: "tenth", gift: "ten Lords-a-Leaping"}
	song[10] = songDay{day: "eleventh", gift: "eleven Pipers Piping"}
	song[11] = songDay{day: "twelfth", gift: "twelve Drummers Drumming"}

	return true
}

// Song will return the entire song
func Song() string {
	initSong()

	var buf bytes.Buffer

	for x := 0; x < 12; x++ {
		buf.Write([]byte("On the "))
		buf.Write([]byte(song[x].day))
		buf.Write([]byte(" day of Christmas my true love gave to me, "))
		buf.Write([]byte(song[x].gift))
		for y := x - 1; y >= 0; y-- {
			buf.Write([]byte(", "))
			if y == 0 {
				buf.Write([]byte("and "))
				buf.Write([]byte(song[y].gift))
				break
			}

			buf.Write([]byte(song[y].gift))
		}
		buf.Write([]byte("\n"))
	}

	return buf.String()
}

// Verse will return a specific portion of the song
func Verse(verse int) string {
	initSong()

	verse = verse - 1 // off-by-one

	var buf bytes.Buffer

	buf.Write([]byte("On the "))
	buf.Write([]byte(song[verse].day))
	buf.Write([]byte(" day of Christmas my true love gave to me, "))
	buf.Write([]byte(song[verse].gift))
	for y := verse - 1; y >= 0; y-- {
		buf.Write([]byte(", "))
		if y == 0 {
			buf.Write([]byte("and "))
			buf.Write([]byte(song[y].gift))
			break
		}

		buf.Write([]byte(song[y].gift))
	}

	return buf.String()
}
