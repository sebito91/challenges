// Package protein converts a given RNA string into its components
package protein

// testVersion is the suite we're up against
const testVersion = 1

// cases is a horrible global variable with our mappings
var cases = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon returns the protein name from a given codon
func FromCodon(input string) string {
	return cases[input]
}

// FromRNA returns the list of proteins in an RNA strand
func FromRNA(input string) (lister []string) {
	for x := 0; x < len(input); x += 3 {
		if cases[input[x:x+3]] == "STOP" {
			return lister
		}

		lister = append(lister, cases[input[x:x+3]])
	}

	return lister
}
