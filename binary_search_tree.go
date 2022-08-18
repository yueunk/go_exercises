package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// adds nodes to the tree (no duplicates)
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k) // recursively call Insert to the Right node
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true // this happens when n.Key == k
}

func (n *Node) traverseAndPrint() {
	if n == nil {
		return
	}
	n.Left.traverseAndPrint()
	fmt.Println(n.Key)
	n.Right.traverseAndPrint()
}

func (n *Node) findMax() int {
	if n.Right == nil {
		return n.Key
	} else {
		return n.Right.findMax()
	}
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(230)
	tree.Insert(350)
	tree.Insert(280)
	tree.Insert(309)
	tree.Insert(206)
	tree.Insert(3009)
	fmt.Println(tree)
	fmt.Println(tree.Search(3009))
	fmt.Println(tree.Search(30099))
	tree.traverseAndPrint()
	fmt.Println(tree.findMax())
}
