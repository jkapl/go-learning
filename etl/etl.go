package etl

import "strings"

// Transform maps letters to their scrabble scores
func Transform(l map[int][]string) map[string]int {
	transformed := make(map[string]int)

	for score, letters := range l {
		for _, letter := range letters {
			transformed[strings.ToLower(letter)] = score
		}
	}

	return transformed
}