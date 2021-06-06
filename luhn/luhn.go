package luhn

import (
	"strings"
	"unicode"
)

// intValid returns true if the input string satisfies the Luhn algorithm
func Valid(i string) bool {
	noSpaces := strings.ReplaceAll(i, " ", "")
	if len(noSpaces) <= 1 {
		return false
	}
	r := []rune(noSpaces)
	sum := 0
	for i, v := range r {
		if !unicode.IsDigit(v) {
			return false
		}
		intV := int(v - '0')
		if (len(r)-1-i)%2 == 1 {
			intV *= 2
			if intV > 9 {
				intV -= 9
			}
			sum += intV
		} else {
			sum += intV
		}
	}
	return sum%10 == 0
}
