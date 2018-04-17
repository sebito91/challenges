// Package house prints out the song The House That Jack Built
package house

import "strings"

const (
	tail = "house that Jack built."
	head = "This is the"
)

// Embed embeds a noun phrase as the object of relative clause with a
// transitive verb.
//
// Argument relPhrase is a phrase with a relative clause, minus the object
// of the clause.  That is, relPhrase consists of a subject, a relative
// pronoun, a transitive verb, possibly a preposition, but then no object.
func Embed(relPhrase, nounPhrase string) string {
	return strings.Join([]string{relPhrase, nounPhrase}, " ")
}

// Verse generates a verse of a song with relative clauses that have
// a recursive structure.
func Verse(subject string, relPhrases []string, nounPhrase string) string {
	relPhrases = append(relPhrases, nounPhrase)
	stuffs := append([]string{subject}, relPhrases...)

	return strings.Join(stuffs, " ")
}

// Song generates the full text of "The House That Jack Built"
func Song() string {
	var output []string

	things := map[int]string{
		1:  "malt",
		2:  "rat",
		3:  "cat",
		4:  "dog",
		5:  "cow with the crumpled horn",
		6:  "maiden all forlorn",
		7:  "man all tattered and torn",
		8:  "priest all shaven and shorn",
		9:  "rooster that crowed in the morn",
		10: "farmer sowing his corn",
		11: "horse and the hound and the horn",
	}

	actions := map[int]string{
		1:  "ate",
		2:  "killed",
		3:  "worried",
		4:  "tossed",
		5:  "milked",
		6:  "kissed",
		7:  "married",
		8:  "woke",
		9:  "kept",
		10: "belonged to",
	}

	output = append(output, Embed(head, tail))

	for count := 1; count < 12; count++ {
		output = append(output, "")
		for each := count; each > 0; each-- {
			if each == count {
				output = append(output, "This is the "+things[each])
				continue
			}
			output = append(output, Embed(Embed("that", actions[each]), Embed("the", things[each])))
		}
		output = append(output, Embed("that lay in the", tail))
	}

	return strings.Join(output, "\n")
}
