package luhn

import (
	"strconv"
	"strings"
	"unicode"
)


// Valid returns true if the input string satisfies the Luhn algorithm
func Valid(i string) bool {
	noSpaces := strings.ReplaceAll(i, " ", "")
	if len(noSpaces) <= 1 {
		return false
	}
	r := []rune(noSpaces)
	sum := 0
	index := 0
	for i := len(r) - 1; i >= 0; i-- {
		if unicode.IsDigit(r[i]) {
			num, _ := strconv.Atoi(string(r[i]))
			if index%2 == 1 {
				num *= 2
				if num > 9 {
					num -= 9
				}
				sum += num
			} else {
				sum += num
			}
			index++
		} else {
			return false
		}
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
