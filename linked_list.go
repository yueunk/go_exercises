package main

import "fmt"

type linkedList struct {
	head   *node
	length int
}

type node struct {
	data int
	next *node
}

// returns a value at the given idx
func (l *linkedList) read(idx int) int {
	// starting at counter 0, once given idx reached, return that val
	counter := 0
	currentNode := l.head

	for counter < l.length {
		if counter == idx {
			return currentNode.data
		}
		counter++
		currentNode = currentNode.next
	}
	return -1
}

// returns the index of the given value
func (l *linkedList) indexOf(val int) int {
	counter := 0
	currentNode := l.head

	for counter < l.length {
		if currentNode.data == val {
			return counter
		}
		counter++
		currentNode = currentNode.next
	}
	return -1
}

// inserts the node with the given value at the given index

// create a new node with the given val
// if inserting at the beginning
// - link the new node to the current head
// - set the head to the new node
// - increment the length by 1
// otherwise
// - set the current node to the node at i - 1
// - link the new node to the current node's next node
// - link the current node to the new node
// - increment the length by 1

func (l *linkedList) insertAtIndex(val, i int) {
	newNode := &node{data: val}

	if i == 0 {
		newNode.next = l.head
		l.head = newNode
		l.length++
		return
	}

	currentNode := l.head
	currentIdx := 0

	for currentIdx < (i - 1) {
		currentNode = currentNode.next
		currentIdx++
	}

	newNode.next = currentNode.next
	currentNode.next = newNode
	l.length++

	return
}

func (l *linkedList) deleteAtIndex(i int) {
	if i == 0 {
		newHead := l.head.next
		l.head.next = nil
		l.head = newHead
		l.length--
		return
	}

	// access the node at i - 1
	// link the node to the next next node
	// set its next node's next to nil
	// decrement by 1
	currentNode := l.head
	currentIdx := 0
	for currentIdx < (i - 1) {
		currentNode = currentNode.next
		currentIdx++
	}
	deletedNode := currentNode.next
	currentNode.next = currentNode.next.next
	deletedNode.next = nil

	l.length--
}

func (l *linkedList) getLastNode() *node {
	// without knowing the length
	// start with the head
	// get to the node that has next node as nil
	// return that node
	currentNode := l.head
	for currentNode != nil {
		if currentNode.next == nil {
			return currentNode
		} else {
			currentNode = currentNode.next
		}
	}
	return currentNode
}

func (l *linkedList) reverseList() {
	var prevNode *node
	currentNode := l.head

	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	l.head = prevNode
}

func main() {
	n3 := &node{data: 3}
	n2 := &node{data: 2, next: n3}
	n1 := &node{data: 1, next: n2}
	myList := &linkedList{head: n1, length: 3}

	fmt.Println(myList.read(0))     // 1
	fmt.Println(myList.read(10))    // -1
	fmt.Println(myList.indexOf(1))  // 0
	fmt.Println(myList.indexOf(10)) // -1

	myList.insertAtIndex(100, 0) // inserts 100 at the beginning
	fmt.Println(myList)          // length should be 4
	fmt.Println(myList.read(0))  // 100
	myList.insertAtIndex(200, 4) // inserts 200 at index 4
	fmt.Println(myList)          // length should be 5
	fmt.Println(myList.read(4))  // 200

	myList.deleteAtIndex(0)     // deletes the first item
	fmt.Println(myList)         // length should be 4
	fmt.Println(myList.read(0)) // 1

	myList.deleteAtIndex(3)     // deletes the first item
	fmt.Println(myList)         // length should be 3
	fmt.Println(myList.read(3)) // -1
	fmt.Println(myList.read(2)) // 3

	fmt.Println(myList.getLastNode()) // node with val 3 and next nil
	myList.reverseList()
	fmt.Println(myList.getLastNode()) // node with val 1 and next nil
}
