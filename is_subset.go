package main

import "fmt"

func isSubset(s1, s2 []int) bool {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1 // s1 is the longer
	}

	hash := map[int]bool{}
	for _, val := range s1 {
		hash[val] = true
	}

	for _, val := range s2 {
		if !hash[val] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isSubset([]int{1, 2, 3}, []int{5, 4, 3, 2, 1})) // true
}
