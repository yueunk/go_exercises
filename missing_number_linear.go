package main

import "fmt"

// O(N) approach
func missingInteger(nums []int) int {
	// get the expected and real sum of the array
	// return the result of subtraction of the real from the expected
	expectedSum := len(nums)
	actualSum := 0
	for idx, num := range nums {
		expectedSum += idx
		actualSum += num
	}
	return expectedSum - actualSum
}

func main() {
	nums1 := []int{2, 3, 0, 6, 1, 5}
	nums2 := []int{8, 2, 3, 9, 4, 7, 5, 0, 6}
	fmt.Println(missingInteger(nums1)) // 4
	fmt.Println(missingInteger(nums2)) // 1
}
