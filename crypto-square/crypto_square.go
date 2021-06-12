package cryptosquare

import ( 
	"unicode"
	// "fmt"
	"math"
)

func Encode(s string) string {
	var normalized []rune
	var final []rune
	for _, val := range s {
		if unicode.IsLetter(val) || unicode.IsNumber(val) {
			normalized = append(normalized, unicode.ToLower(val))
		}
	}
	var sq = math.Sqrt(float64(len(s)))
	var c int
	// var r int
	var diff = sq - math.Round(sq)
	if diff != 0 {
		if diff > 0.5 {
			c = int(math.Round(sq))
			// r = int(math.Round(sq) - 1)
		} else {
			c = int(math.Round(sq) + 1)
			// r = int(math.Round(sq))
		}
	} else {
		c = int(sq)
		// r = int(sq)
	} 
	var ptr = 0
	var col = 0
	// var row = 0
	var totalLength = len(normalized)
	for totalLength % c != 0 {
		totalLength += 1
	}
	for i := 0; i < totalLength; i++ {
		// if i > len(normalized) {
		// 	final = append(final, ' ')
			
		// }
		if ptr >= len(normalized) {
			col += 1
			ptr = col
			final = append(final, ' ')
		}
		final = append(final, normalized[ptr])
		ptr += c
	}
	return string(final)
}