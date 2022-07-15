package main

import "fmt"

func insertionSort(s []int) []int {
	for i := 1; i < len(s); i += 1 {
		val := s[i]
		prev := i - 1

		for prev >= 0 {
			if val < s[prev] {
				s[prev+1] = s[prev]
				prev -= 1
			} else {
				break
			}
		}
		s[prev+1] = val
		fmt.Println(s)
	}

	return s
}

func main() {
	slc := []int{65, 55, 45, 35, 25, 150, 10}
	fmt.Println(insertionSort(slc)) // [10 25 35 45 55 65 150]
}
