package main

import (
	"fmt"
	"sort"
)

func findMissingNumber(arr []int) int {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{5, 1, 2, 4, 0}
	fmt.Println(findMissingNumber(nums)) // 3
}
