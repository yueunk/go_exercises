package main

import "fmt"

func addUntil100(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	nextSum := addUntil100(arr[1:])

	if arr[0]+nextSum > 100 {
		return nextSum
	} else {
		return arr[0] + nextSum
	}
}

func main() {
	fmt.Println(addUntil100([]int{130, 2, 3, 30})) // 35
}
