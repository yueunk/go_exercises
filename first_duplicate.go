package main

import "fmt"

// input: arr of strings; output: str
// given an array of strings, find the first duplicate value at O(N)
// - iterate through the arr of strings
// 	- convert it to a map w/ true vals
//  - if the current val already exists as a key, return the val

func firstDuplicate(s []string) string {
	hash := map[string]bool{}
	for _, val := range s {
		if hash[val] {
			return val
		}
		hash[val] = true
	}
	return ""
}

func main() {
	fmt.Println(firstDuplicate([]string{"a", "b", "c", "d", "c", "e", "f"}))
}
