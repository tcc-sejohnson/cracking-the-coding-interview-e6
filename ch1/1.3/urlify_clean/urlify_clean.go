package urlify_clean

import "strings"

// This is one of those problems that I hate. It's a solved problem, so
// here's what I'd consider the "idiomatic" way to do it. I'll write another
// library-less version as well.
func Urlify(input string, length int) string {
	// We receive a string that's padded to "hold" the extra characters
	input = strings.Trim(input, " ")
	input = strings.ReplaceAll(input, " ", "%20")
	// We're making the assumption that the provided length is correct
	// but this could be easily checked at this point.
	return input
}
