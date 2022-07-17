package main

import "fmt"

// input: arr of int; output: new arr w/ even nums
// base case: arr len 0 ==> return []
// sub problem
// - evenNumbers(arr[:i-1]) returns an arr of even numbers so far
// - if arr[i] is even, add to the result

func evenNumbers(arr []int) []int {
	result := []int{}
	if len(arr) == 0 {
		return result
	}
	i := len(arr) - 1
	result = evenNumbers(arr[:i-1])
	if arr[i]%2 == 0 {
		result = append(result, arr[i])
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(evenNumbers(arr)) // [2 4 6 8 10]
}
