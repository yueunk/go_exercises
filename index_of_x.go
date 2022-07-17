package main

import "fmt"

/*
input: string, output: first idx that contains "x"
assumption: the given string always has at least an "x"

base case: if str[0] is x, return 0
sub problem ("bxyz" of "abxyz")
- indexOfX(str[1:]) returns an index of x as 1
- add 1 to that index
*/

func indexOfX(str string) int {
	if str[0] == 'x' {
		return 0
	}
	return indexOfX(str[1:]) + 1
}

func main() {
	str1 := "abcdefghijklmnopqrstuvwxyz"
	str2 := "axcdefghijklmnopqrstuvwxyz"
	fmt.Println(indexOfX(str1)) // 23
	fmt.Println(indexOfX(str2)) // 1
}
