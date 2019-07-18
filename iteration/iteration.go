package iteration

import "strings"

// Repeat takes a string containing a character and repeats it 5 times
func Repeat(character string, iterations int) string {
	return strings.Repeat(character, iterations)
}
