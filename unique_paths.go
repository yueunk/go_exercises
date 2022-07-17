package main

import "fmt"

/*
input: num rows, num columns
output: num of shortest paths from the topleft to bottomright
base case: if row or col 1, return 1
*/

func uniquePaths(row, col int) int {
	if row == 1 || col == 1 {
		return 1
	}

	return uniquePaths(row-1, col) + uniquePaths(row, col-1)
}

func main() {
	fmt.Println(uniquePaths(3, 4)) // 10
}
