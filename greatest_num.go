package main

import "fmt"

// get the greatest num in linear time
// assume the first elem is the greatest, compare w/ the rest and reset greatest accordingly
func greatestNum(s []int) int {
	greatest := s[0]
	for _, val := range s[1:] {
		if val > greatest {
			greatest = val
		}
	}

	return greatest
}

func main() {
	slc := []int{65, 55, 45, 35, 25, 150, 10}
	fmt.Println(greatestNum(slc)) // 150
}
