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
	sum := 0
	even := len(noSpaces)%2 == 0
	for _, v := range noSpaces {
		if !unicode.IsDigit(v) {
			return false
		}
		intV := int(v - '0')
		if even {
			intV *= 2
			if intV > 9 {
				intV -= 9
			}
			sum += intV
		} else {
			sum += intV
		}
		even = !even
	}
	return sum%10 == 0
}
