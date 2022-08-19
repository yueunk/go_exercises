package main

import "fmt"

// assumption: arr has at least 1 positive number; subsection is contiguous
func maxSum(arr []int) int {
	currentSum := 0
	greatestSum := 0

	for _, val := range arr {
		if currentSum+val < 0 {
			currentSum = 0
		} else {
			currentSum += val
			if currentSum > greatestSum {
				greatestSum = currentSum
			}
		}
	}

	return greatestSum
}

func main() {
	arr1 := []int{3, -4, 4, -3, 5, -9}
	arr2 := []int{-4, -4, -3, -5, 9}
	fmt.Println(maxSum(arr1)) // 6
	fmt.Println(maxSum(arr2)) // 9
}
