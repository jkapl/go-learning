package strain

type Ints []int

type Lists [][]int

type Strings []string

func (i Ints) Keep (f func ( int) bool) Ints {
	res := Ints{}
	for _,val := range i {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

func (i Ints) Discard (func (int) bool) Ints {
	return nil
} 

func (l Lists) Keep (func ([]int) bool) Lists {
	return nil
}

func (s Strings) Keep (func(string) bool) Strings {
	return nil
}