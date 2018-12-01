package isogram

import "unicode"

// IsIsogram determines if a word is or is not an isogram
// An isogram is a word with no repeating letters
// Spaces and dashes _are_ allowed to be repeated
func IsIsogram(word string) bool {
	foundLetters := map[rune]bool{}
	for _, char := range word {
		charUpper := unicode.ToUpper(char)
		if !unicode.IsLetter(charUpper) {
			continue
		}
		if _, ok := foundLetters[charUpper]; ok {
			return false
		}
		foundLetters[charUpper] = true
	}
	return true
}
