package isogram

import "unicode"

// IsIsogram checks if a string is an isogram
func IsIsogram(s string) bool {
	m := make(map[rune]bool)
	for _, val := range s {
		if !(unicode.IsLetter(val)) {
			continue
		}
		val = unicode.ToLower(val)
		if m[val] {
			return false
		}
		m[val] = true
	}
	return true
}
