package luhn

import (
	"strings"
	"unicode"
)

// Valid returns true if the id string satisfies the Luhn algorithm
func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	sum := 0
	even := len(id)%2 == 0
	for _, v := range id {
		if !unicode.IsDigit(v) {
			return false
		}
		digit := int(v - '0')
		if even {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		even = !even
	}
	return sum%10 == 0
}
