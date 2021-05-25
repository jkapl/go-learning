package accumulate

func Accumulate(input []string, acc func(string) string) []string {
	accResult := []string{}
	if (len(input) == 0) {
		return input
	}
	for _,val := range input {
		changed := acc(val)
		accResult = append(accResult, changed)
	}
	return accResult
}
