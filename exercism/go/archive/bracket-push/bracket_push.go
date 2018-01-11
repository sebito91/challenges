// Package brackets is a matcher for open/closed parens, etc
package brackets

// testVersion is the suite we're up against
const testVersion = 4

var lister = map[string]string{
	"{": "}",
	"[": "]",
	"(": ")",
}

// Bracket is our bracket/paren matcher
func Bracket(input string) (works bool, err error) {
	status := []string{}

	for _, each := range input {
		if _, ok := lister[string(each)]; ok {
			status = append(status, string(each))
		} else {
			if len(status) > 0 && lister[status[len(status)-1]] == string(each) {
				status = status[:len(status)-1]
			} else {
				return false, nil
			}
		}
	}

	if len(status) > 0 {
		return false, nil
	}

	return true, nil
}
