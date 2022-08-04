package main

import (
	"fmt"
	"sort"
)

func greatestProduct(nums []int) int {
	sort.Ints(nums)
	idx := len(nums) - 3
	return nums[idx] * nums[idx+1] * nums[idx+2]

}

func main() {
	nums := []int{3, 1, 2, 4, 6}
	fmt.Println(greatestProduct(nums)) // 72
}
