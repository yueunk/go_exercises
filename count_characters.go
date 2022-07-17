package main

import "fmt"

// input: arr of strings, output: int (count of all chrs)
// basecase: if arr length 0, return 0
// sub-problem:
// - countCharacters(arr[1:]) returns the total count so far
// - add the first string's count to the total count above
func countCharacters(arr []string) int {
	if len(arr) == 0 {
		return 0
	}
	return len(arr[0]) + countCharacters(arr[1:])
}

func main() {
	arr := []string{"ab", "c", "def", "ghij"}
	fmt.Println(countCharacters(arr)) // 10
}
