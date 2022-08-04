package main

import (
	"fmt"
	"sort"
)

func hasDuplicateValue(nums []int) bool {
	sort.Ints(nums)
	// iterate through the slice (idx 0 to slice size - 2)
	// if the current val is the same as the next val, return true; otherwise return false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{3, 1, 2, 4, 6}
	nums2 := []int{3, 3, 2, 4, 6}
	fmt.Println(hasDuplicateValue(nums))  // false
	fmt.Println(hasDuplicateValue(nums2)) // true
}
