package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns a string based on factors of the input integer
func Convert(num int) string {
	var sb strings.Builder
	if num%3 == 0 {
		sb.WriteString("Pling")
	}
	if num%5 == 0 {
		sb.WriteString("Plang")
	}
	if num%7 == 0 {
		sb.WriteString("Plong")
	}
	if len(sb.String()) == 0 {
		return strconv.Itoa(num)
	}
	return sb.String()
}
