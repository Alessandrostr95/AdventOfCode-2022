package main

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}

	e := (*s)[s.Len()-1]

	if s.Len() == 1 {
		*s = make([]int, 0)
	} else {
		*s = (*s)[:s.Len()-1]
	}

	return e, true
}

func (s *Stack) Top() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}
	return (*s)[s.Len()-1], true
}
