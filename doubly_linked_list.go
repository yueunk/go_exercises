package main

import "fmt"

type doublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	data int
	prev *node
	next *node
}

func (l *doublyLinkedList) insertAtEnd(val int) {
	newNode := &node{data: val}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length = 1
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
		l.length++
	}
}

func (l *doublyLinkedList) removeFromFront() *node {
	removedNode := l.head
	l.head = removedNode.next
	l.head.prev = nil
	l.length--
	return removedNode
}

func (l *doublyLinkedList) printAll() {
	currentNode := l.head
	fmt.Printf("%s", "all values: ")
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("")
}

func (l *doublyLinkedList) printAllReverse() {
	currentNode := l.tail
	fmt.Printf("%s", "all values in reverse: ")
	for i := l.length - 1; i >= 0; i-- {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.prev
	}
	fmt.Println("")
}

func main() {
	n1 := &node{data: 1}
	n2 := &node{data: 2}
	n3 := &node{data: 3}
	n1.next = n2
	n2.prev = n1
	n2.next = n3
	n3.prev = n2
	myList := &doublyLinkedList{}
	myList.head = n1
	myList.tail = n3
	myList.length = 3

	fmt.Println(myList) // length 3 with both head and tail
	myList.printAll()   // all values: 1 2 3

	myList.insertAtEnd(4)
	fmt.Println(myList.tail) // a new node inserted with value 4 and nil next
	fmt.Println(myList)      // length 4 with new tail
	myList.printAll()        // all values: 1 2 3 4

	fmt.Println(myList.removeFromFront()) // removed node with value 1
	fmt.Println(myList)                   // length 3 with new head
	myList.printAll()                     // all values: 2 3 4
	myList.printAllReverse()              // all values in reverse: 4 3 2
}
