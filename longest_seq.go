package main

import "fmt"

func longestSequence(arr []int) int {
	maxStreak := 1
	currentStreak := 1
	hash := map[int]bool{}
	min := arr[0]
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		hash[arr[i]] = true
		if arr[i] <= min {
			min = arr[i]
		} else if arr[i] >= max {
			max = arr[i]
		}
	}

	for j := min; j < max; j++ {
		if hash[j] && hash[j+1] {
			currentStreak++
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		} else {
			currentStreak = 1
		}
	}

	return maxStreak
}

func main() {
	arr1 := []int{10, 5, 12, 3, 55, 30, 4, 11, 2}
	arr2 := []int{19, 13, 15, 12, 18, 14, 17, 11}
	fmt.Println(longestSequence(arr1)) // 4
	fmt.Println(longestSequence(arr2)) // 5
}
