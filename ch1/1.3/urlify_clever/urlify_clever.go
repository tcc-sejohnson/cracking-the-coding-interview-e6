package urlify_clever

// Okay... if we're really not allowed to do this the "right" way, we
// can at least do it the "clever" way, using no libraries.
func Urlify(input string, length int) string {
	if length <= 0 {
		return ""
	}
	runes := []rune(input)
	fillIndex := len(runes) - 1
	for i := length - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			runes[fillIndex-2] = '%'
			runes[fillIndex-1] = '2'
			runes[fillIndex] = '0'
			fillIndex -= 3
		} else {
			runes[fillIndex] = runes[i]
			fillIndex -= 1
		}
	}
	return string(runes)
}
