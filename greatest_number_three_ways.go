package main

import (
	"fmt"
	"sort"
)

// O(N log N) -- built-in sort
func greatestNumber1(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

// O(N)
func greatestNumber2(arr []int) int {
	greatest := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > greatest {
			greatest = arr[i]
		}
	}
	return greatest
}

// O(N^2) -- bubble sort
// compare two vals, if the second val smaller, swap
// continue til no more swaps
func greatestNumber3(arr []int) int {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}
	return arr[len(arr)-1]
}

func main() {
	nums := []int{9, 0, 2, 8, 1}
	fmt.Println(greatestNumber1(nums)) // 9
	fmt.Println(greatestNumber2(nums)) // 9
	fmt.Println(greatestNumber3(nums)) // 9
}
