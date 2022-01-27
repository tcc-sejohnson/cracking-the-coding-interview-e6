package check_permutation_revised

// After looking at some hints, sorting seems the easiest
// to understand and most obvious. However, a good sort has O(log(n)) time
// complexity. I think we can get this down to O(n) time by using a map.

func IsPermutation(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	firstRuneCounts := make(map[rune]int)
	for _, r := range first {
		firstRuneCounts[r] += 1
	}

	for _, r := range second {
		firstRuneCounts[r] -= 1
		// We want to be able to exit early if possible
		if firstRuneCounts[r] < 0 {
			return false
		}
	}

	return true
}
