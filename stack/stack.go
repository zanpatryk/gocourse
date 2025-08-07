package stack

import "fmt"

type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() rune {
	if s.isEmpty() {
		return 0
	}
	lastChar := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return lastChar
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s Stack) Print() {
	fmt.Println(string(s.items))
}
