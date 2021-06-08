package grains

func Total() uint64 {
	var sum uint64
	sum = 1
	for i:=1;i<=64;i++{
		sum*=2
	}
	return sum
}

func Square(sq int) (uint64, error) {
	var value uint64
	value = 1
	return value, nil
}