package main

import "fmt"

func sumSwap(a1, a2 []int) []int {
	hash := map[int]int{}
	sum1 := 0
	sum2 := 0

	for idx, val := range a1 {
		sum1 += val
		hash[val] = idx
	}

	for _, val := range a2 {
		sum2 += val
	}

	shiftAmount := (sum1 - sum2) / 2

	for idx, val := range a2 {
		_, ok := hash[val+shiftAmount]
		if ok {
			return []int{hash[val+shiftAmount], idx}
		}
	}

	return []int{-1, -1}
}

func main() {
	arr1 := []int{5, 3, 3, 7}
	arr2 := []int{4, 1, 1, 6}
	fmt.Println(sumSwap(arr1, arr2)) // [3, 0]
}
