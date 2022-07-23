package main

import "fmt"

// memo[n] references golomb(n)
func golomb(n int, memo map[int]int) int {
	if n == 1 {
		return n
	}

	if memo[n] == 0 {
		memo[n] = 1 + golomb(n-golomb(golomb(n-1, memo), memo), memo)
	}

	return memo[n]
}

func main() {
	memo := map[int]int{}
	fmt.Println(golomb(12, memo)) // 6
	fmt.Println(golomb(15, memo)) // 6
	fmt.Println(golomb(20, memo)) // 8
}
