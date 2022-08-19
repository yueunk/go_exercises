package main

import "fmt"

func max(arr []int) int {
	// assume the first elem as max
	// - iterate through the arr starting the second val
	// - replace the max if current val bigger than the max
	// return max
	greatest := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > greatest {
			greatest = arr[i]
		}
	}
	return greatest
}

func main() {
	arr1 := []int{5, 3, 7}
	arr2 := []int{4, 1, 60}
	fmt.Println(max(arr1)) // 7
	fmt.Println(max(arr2)) // 60
}
