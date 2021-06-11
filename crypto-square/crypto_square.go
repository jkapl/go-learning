package cryptosquare

import ( 
	"unicode"
	"fmt"
	"math"
)

func Encode(s string) string {
	var normalized []rune
	var final []rune
	for _, val := range s {
		if unicode.IsLetter(val) {
			normalized = append(normalized, unicode.ToLower(val))
		}
	}
	fmt.Println(normalized)
	var sq = math.Sqrt(float64(len(s)))
	var c float64
	var r float64
	var diff = sq - math.Round(sq)
	if diff != 0 {
		if diff > 0.5 {
			c = math.Round(sq)
			r = math.Round(sq) - 1
		} else {
			c = math.Round(sq) + 1
			r = math.Round(sq)
		}
	} else {
		c = sq
		r = sq
	} 
	for i, val := range normalized {
			
	}


}