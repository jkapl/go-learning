package luhn

import ( 
	"strings"
	"strconv"
	"unicode"
)

func Valid (i string) bool {
	if len(i) <= 1 {
		return false
	}
	noSpaces := strings.ReplaceAll(i, " ", "")
	sum := 0
	for i, v := range noSpaces {
		if (unicode.IsDigit(v)) {
			num, _ := strconv.Atoi(string(v))
			if i % 2 == 0 {
				num *= 2
				if (num > 9) {
					num -= 9
				}
			}
			sum += num
		}
		return false
	}
	if sum % 10 == 0 {
		return true
	}
	return false
}
