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

func (s *Stack) Read() rune {
	return s.items[len(s.items)-1]
}

func lint(str string) bool {
	lintStack := Stack{}
	match := map[rune]rune{')': '(', ']': '[', '}': '{'}
	opening := map[rune]bool{'(': true, '[': true, '{': true}
	closing := map[rune]bool{')': true, ']': true, '}': true}

	for _, r := range str {
		if !opening[r] && !closing[r] {
			continue
		}

		if opening[r] {
			lintStack.Push(r)
		} else if closing[r] && lintStack.Read() == match[r] {
			lintStack.Pop()
		} else {
			return false
		}
	}

	if len(lintStack.items) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(lint("(var x = {y: [1, 2, 3])}")) // false
	fmt.Println(lint("(var x = {y: [1, 2, 3]})")) // true
}
