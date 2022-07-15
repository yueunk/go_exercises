package main

import "fmt"

func factorial(n int) int {
	// base case: if 1 given, return 1
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(10))
}
