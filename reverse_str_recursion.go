package main

import "fmt"

func reverse(str string) string {
	if len(str) == 1 {
		return str
	}
	return reverse(str[1:]) + string(str[0])
}

func main() {
  fmt.Println(reverse("12345")) // 54321
}
