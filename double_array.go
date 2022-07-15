package main

import "fmt"

func doubleArray(arr []int, idx int) {
	if idx >= len(arr) {
		return
	}
	arr[idx] *= 2
	doubleArray(arr, idx+1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	doubleArray(arr, 0)
	fmt.Println(arr) // [2, 4, 6, 8, 10]
}
