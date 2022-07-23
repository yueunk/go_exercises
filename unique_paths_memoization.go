package main

import "fmt"

func uniquePaths(r, c int, memo map[[2]int]int) int {
	if r == 1 || c == 1 {
		return 1
	}

	arr := [2]int{r, c}
	if memo[arr] == 0 {
		memo[arr] = uniquePaths(r-1, c, memo) + uniquePaths(r, c-1, memo)
	}
	return memo[arr]
}

func main() {
	memo := map[[2]int]int{}
	fmt.Println(uniquePaths(1, 5, memo)) // 1
	fmt.Println(uniquePaths(4, 3, memo)) // 10
}
