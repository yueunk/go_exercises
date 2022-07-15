package main

import "fmt"

// bubble sort in Go
// - while toggle false
// - iterate through the slice (idx 0 to length - 2)
// - compare the elem at current idx with the next
//   - if current bigger, swap
//   - toggle swap to true

func bubbleSort(s []int) []int {
	upper := len(s) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < upper; i += 1 {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				sorted = false
			}
		}
		upper -= 1
	}

	return a
}

func main() {
	slc := []int{65, 55, 45, 35, 25, 15, 10}
	fmt.Println(bubbleSort(slc)) // [10 15 25 35 45 55 65]
}
