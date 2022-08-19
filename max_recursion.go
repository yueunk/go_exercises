package main

import "fmt"

func max(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	nextMax := max(arr[1:])
	if arr[0] > nextMax {
		return arr[0]
	} else {
		return nextMax
	}
}

func main() {
	fmt.Println(max([]int{10, 2, 3})) // 10
}
