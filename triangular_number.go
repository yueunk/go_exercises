package main

import "fmt"

/*
compute triagular num given the 1-based position
- triagular nums: 1, 3, 6, 10, ..., num; num is the sum of the prev num & the current position num

input: 1-based index "n"; output: triagular num at the nth position
base case: if the n is 0, return 0
sub problem:
- triagularNumber(n - 1) returns its triagularNumber
- add n to the return value
*/

func triagularNumber(n int) int {
	if n <= 0 {
		return 0
	}

	return triagularNumber(n-1) + n
}

func main() {
	fmt.Println(triagularNumber(7)) // 28
}
