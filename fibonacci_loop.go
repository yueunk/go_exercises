package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}

	a := 0
	b := 1

	for i := 1; i < n; i += 1 {
		temp := a
		a = b
		b = temp + a
	}

	return b
}

func main() {
	fmt.Println(fib(10)) // 55
}
