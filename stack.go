package main

import (
	"fmt"
)

// Stack represents a stack that holds a slice; LIFO
type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// last out
func (s *Stack) Pop() int {
	l := len(s.items)
	removed := s.items[l-1]
	s.items = s.items[:l-1]
	return removed
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack) // {[]}
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack) // {[1 2 3]}
	myStack.Pop()
	myStack.Pop()
	fmt.Println(myStack) //{[1]}
}
