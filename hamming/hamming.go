package hamming

import "errors"

// Distance returns the Hamming distance between two strings
// It returns the distance as an int or an error
func Distance(a, b string) (int, error) {
	dist := 0
	bLen := len(b)
	for idx, letter := range a {
		if (idx + 1) > bLen {
			dist++
		} else {
			if string(letter) != string(b[idx]) {
				dist++
			}
		}
	}
	if len(b) != len(a) {
		return 0, errors.New("Length erro")
	}
	return dist, nil
}
