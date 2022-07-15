package main

import (
	"fmt"
)

// Queue represents a queue that holds a slice; FIFO
type Queue struct {
	items []int
}

func (s *Queue) Enqueue(i int) {
	s.items = append(s.items, i)
}

// first out
func (s *Queue) Dequeue() int {
	removed := s.items[0]
	s.items = s.items[1:]
	return removed
}

func main() {
	myStack := Queue{}
	fmt.Println(myStack) // {[]}
	myStack.Enqueue(1)
	myStack.Enqueue(2)
	myStack.Enqueue(3)
	fmt.Println(myStack) // {[1 2 3]}
	myStack.Dequeue()
	myStack.Dequeue()
	fmt.Println(myStack) //{[3]}
}
