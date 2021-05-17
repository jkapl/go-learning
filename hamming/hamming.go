package hamming

import "errors"

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
	} else {
		return dist, nil
	}
}
