package main

import "fmt"

type linkedList struct {
	head *node
}

type node struct {
	data int
	next *node
}

func (l *linkedList) getAllValues() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func removeMiddleNode(n *node) {
	n.data = n.next.data
	n.next = n.next.next
}

func main() {
	n3 := &node{data: 3}
	n2 := &node{data: 2, next: n3}
	n1 := &node{data: 1, next: n2}
	myList := &linkedList{head: n1}

	myList.getAllValues() // 1 2 3
	removeMiddleNode(n2)
	myList.getAllValues() // 1 3
}
