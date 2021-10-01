package main

import "fmt"

type Stack struct {
	St []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(num int) {
	s.St = append(s.St, num)
}

func (s *Stack) Pop() int {
	ln := len(s.St)
	if ln != 0 {
		ln -= 1
		num := s.St[ln]
		s.St = s.St[:ln]
		return num
	}
	return 0
}

func main() {
	st := New()
	for i := 0; i < 10; i++ {
		st.Push(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", st.Pop())
	}
}
