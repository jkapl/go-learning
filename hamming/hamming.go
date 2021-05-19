package hamming

import "errors"

// Distance returns the Hamming distance between two strings
// It returns the distance as an int or an error
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(br) != len(ar) {
		return 0, errors.New("Strings are of different lengths")
	}
	dist := 0
	for idx := range ar {
		if ar[idx] != br[idx] {
			dist++
		}
	}
	return dist, nil
}
