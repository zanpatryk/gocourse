package stack

import "fmt"

/*
Basic stack implementation
*/

type Stack struct {
	items []rune
}

/*
Pushes a new element to the stack
*/
func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

/*
Pops the last element from the stack
*/
func (s *Stack) Pop() rune {
	if s.isEmpty() {
		return 0
	}
	lastChar := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return lastChar
}

/*
Checks if the stack is empty
*/
func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

/*
Prints the stack
*/
func (s Stack) Print() {
	fmt.Println(string(s.items))
}
