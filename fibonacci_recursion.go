package main

import "fmt"

func fib(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fib(n-2, memo) + fib(n-1, memo)
	return memo[n]
}

func main() {
	fmt.Println(fib(10, map[int]int{})) // 55
}
