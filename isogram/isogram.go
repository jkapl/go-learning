package isogram

import "unicode"

func IsIsogram (s string) bool {
	m := make(map[rune]bool)
	for _, val := range s {
		if val == '-' || val == ' ' {
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
