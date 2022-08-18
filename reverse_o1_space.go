package main

import "fmt"

func reverse(arr []int) []int {
	// reverse the array in place
	// - get midpoint index
	// - iterate through the array up to the midpoint
	// - swap the front and back vals
	// - return the arr
	midpoint := len(arr) / 2
	for i := 0; i <= midpoint; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(reverse(arr1))
	fmt.Println(reverse(arr2))
}
