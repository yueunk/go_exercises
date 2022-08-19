package main

import "fmt"

func twoSum(arr []int) bool {
	// convert the arr into hash
	// iterate through the array
	// - if 10 - the current elem's value is true, return true
	// return false

	hash := map[int]bool{}

	for i := 0; i < len(arr); i++ {
		if hash[10-arr[i]] {
			return true
		}
		hash[arr[i]] = true
	}

	return false
}

func main() {
	arr1 := []int{2, 0, 4, 1, 7, 9}
	arr2 := []int{2, 0, 3, 1}
	fmt.Println(twoSum(arr1)) // true
	fmt.Println(twoSum(arr2)) // false
}
