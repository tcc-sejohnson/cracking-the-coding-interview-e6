package check_permutation

import "sort"

// My initial thought is that if we sort both strings, we can simply compare them.
// I'm also assuming that all UTF-8 runes are acceptable. If they weren't, we'd
// have to trim non-allowed runes out or return an error.
// We're also assuming this is NOT supposed to be case-insensitive.
func IsPermutation(first, second string) bool {
	if len(first) != len(second) {
		return false // different lengths can't be permutations
	}

	// Make our strings mutable to save memory and avoid comparing bytes instead of runes
	firstR := []rune(first)
	secondR := []rune(second)
	sortRunes(firstR)
	sortRunes(secondR)

	for i := range firstR {
		if firstR[i] != secondR[i] {
			return false
		}
	}
	return true
}

func sortRunes(r []rune) {
	sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
}
