package main

import "fmt"

func selectionSort(s []int) []int {
	for i := 0; i < len(s)-1; i += 1 {
		lowest := s[i]
		lowestIdx := i
		for j, val := range s[i+1:] {
			if val < lowest {
				lowest = val
				lowestIdx = j + i + 1
			}
		}
		fmt.Println(lowestIdx)
		s[lowestIdx], s[i] = s[i], s[lowestIdx]
	}

	return s
}

func main() {
	slc := []int{65, 55, 45, 35, 25, 150, 10}
	fmt.Println(selectionSort(slc)) // [10 25 35 45 55 65 150]
}
