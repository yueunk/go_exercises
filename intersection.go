package main

import "fmt"

// find common values from two given arrays and return those values in a new array at constant time
// - create a map and result slice
// - iterate through the first given slice
//	- create a k-v pair with the current item and true
// - iterate through the second given slice
//  - if the current val exists in the map, push to the result slice
// - return result

func intersection(s1, s2 []int) []int {
	result := []int{}
	hash := map[int]bool{}
	for _, val := range s1 {
		hash[val] = true
	}

	for _, val := range s2 {
		if hash[val] {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8})) // [2, 4]
}
