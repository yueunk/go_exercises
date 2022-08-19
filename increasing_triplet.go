package main

import (
	"fmt"
	"math"
)

func increasingTriplet(arr []float64) bool {
	lowest := arr[0]
	middle := math.Inf(1)

	for i := 1; i < len(arr); i++ {
		if arr[i] <= lowest {
			lowest = arr[i]
		} else if arr[i] <= middle {
			middle = arr[i]
		} else {
			return true
		}
	}
	return false
}

func main() {
	arr1 := []float64{5, 2, 8, 4, 3, 7}
	arr2 := []float64{8, 9, 7, 10}
	fmt.Println(increasingTriplet(arr1)) // true
	fmt.Println(increasingTriplet(arr2)) // true
}
