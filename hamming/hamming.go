package hamming

import "errors"

// Distance returns the Hamming distance between two strings
// It returns the distance as an int or an error
func Distance(a, b string) (int, error) {
	if len(b) != len(a) {
		return 0, errors.New("Strings are of different lengths")
	}
	dist := 0
	for idx, letter := range a {
		if string(letter) != string(b[idx]) {
			dist++
		}
	}
	return dist, nil
}
