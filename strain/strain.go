package strain

type Ints struct {
	X []int
}

type Lists struct {
	X [][]int
}

type Strings struct {
	X []string
}

func (i Ints) Keep (f func (x int) bool) Ints {
	r := Ints{}
	if f(x) {
		
	}
}

func (i Ints) Discard (func (int) bool) Ints {

} 

func (l Lists) Keep (func ([]int) bool) Lists {

}

func (s Strings) Keep (func(string) bool) Strings {

}