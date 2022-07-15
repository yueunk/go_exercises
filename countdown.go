package main

import "fmt"

func countdown(i int) {
	if i < 0 {
		return
	}

	fmt.Println(i)

	countdown(i - 1)
}

func main() {
	countdown(10)
}
