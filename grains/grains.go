package grains

import "errors"

// Total returns 2^64 using bit manipulation
func Total() uint64 {
	return 1<<64 - 1
}

// Square returns 2^n where n is the input
func Square(sq int) (uint64, error) {
	// var value uint64 = 1
	if sq <= 0 || sq > 64 {
		return 0, errors.New("Error")
	}
	return 1 << (sq - 1), nil
}
