package is_unique

func IsUnique(input string) bool {
	if input == "" {
		return false
	}
	runeMap := make(map[rune]struct{})
	for _, r := range input {
		if _, exists := runeMap[r]; exists {
			return false
		}
		runeMap[r] = struct{}{}
	}
	return true
}
