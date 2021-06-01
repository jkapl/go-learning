package strain

// Ints is an array of integers
type Ints []int

// Lists is an array of arrays of integers
type Lists [][]int

// Strings is an array of strings
type Strings []string

// Keep is a method of Ints and accepts a function
// It returns the result of calling that function
// on each member of the instance referenced
func (i *Ints) Keep(f func(int) bool) Ints {
	res := Ints{}
	if *i == nil {
		return nil
	}
	for _, val := range *i {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

// Discard is a method of Ints and accepts a function
// It returns the result of calling that function
// on each member of the instance referenced
func (i *Ints) Discard(f func(int) bool) Ints {
	res := Ints{}
	if *i == nil {
		return nil
	}
	for _, val := range *i {
		if !f(val) {
			res = append(res, val)
		}
	}
	return res
}

// Keep is a method of Strings and accepts a function
// It returns the result of calling that function
// on each member of the instance referenced
func (s *Strings) Keep(f func(string) bool) Strings {
	res := Strings{}
	for _, val := range *s {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

// Keep is a method of Lists and accepts a function
// It returns the result of calling that function
// on each member of the instance referenced
func (l *Lists) Keep(f func([]int) bool) Lists {
	re := Lists{}
	for _, val := range *l {
		if f(val) {
			re = append(re, val)
		}
	}
	return re
}
