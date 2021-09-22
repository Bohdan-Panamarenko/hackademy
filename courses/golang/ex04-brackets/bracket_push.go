package main

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

func Bracket(str string) (bool, error) {
	stack := New()
	openBrackets := map[string]int{ "{" : 1, "[" : 2, "(" : 3 }
	closeBrackets := map[string]int{ "}" : 1, "]" : 2, ")" : 3 }
	for _, value := range str {
		open := openBrackets[string(value)]
		close := closeBrackets[string(value)]
		if open != 0 {
			stack.Push(open)
		} else if close != 0 {
			if (stack.Pop() != close) {
				return false, nil
			}
		}
	}
	if (stack.Pop() != 0) {
		return false, nil
	}
	return true, nil
}