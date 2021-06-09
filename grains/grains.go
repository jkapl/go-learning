package grains

import "errors"

func Total() uint64 {
	var sum uint64 = 1
	// sum = 1
	for i := 1; i <= 64; i++ {
		// sum*=2
		sum = sum << 1
		// fmt.Println(sum)
	}
	// fmt.Println(sum)
	return sum - 1
}

func Square(sq int) (uint64, error) {
	var value uint64 = 1
	if sq <= 0 || sq > 64 {
		return 0, errors.New("Error")
	}
	return value << uint64(sq-1), nil
}
