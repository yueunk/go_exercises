package main

import "fmt"

func countX(str string) int {
	if len(str) == 0 {
		return 0
	}

	if str[0] == 'x' {
		return 1 + countX(str[1:])
	} else {
		return countX(str[1:])
	}
}

func main() {
	fmt.Println(countX("axbxcxd")) // 3
}
