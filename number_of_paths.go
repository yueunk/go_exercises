package main

import "fmt"

func numberOfPaths(n int) int {
	if n < 0 {
		return 0
	} else if n < 2 {
		return 1
	}
	return numberOfPaths(n-1) + numberOfPaths(n-2) + numberOfPaths(n-3)
}

func main() {
	fmt.Println(numberOfPaths(11)) // 504
}
