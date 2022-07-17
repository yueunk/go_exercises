package main

import (
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(r rune) {
	s.items = append(s.items, r)
}

func (s *Stack) Pop() rune {
	l := len(s.items)
	removed := s.items[l-1]
	s.items = s.items[:l-1]
	return removed
}

func reverse(str string) string {
	stack := Stack{}
	size := len(str)
	result := ""

	for _, r := range str {
		stack.Push(r)
	}

	for i := 0; i < size; i++ {
		result += string([]rune{stack.Pop()})
	}

	return result
}

func main() {
	fmt.Println(reverse("abcde")) // edcba
}
