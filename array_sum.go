package main

import "fmt"

func arraySum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + arraySum(arr[1:])
}

func main() {
	fmt.Println(arraySum([]int{1, 2, 3, 4, 5})) // 15
}
