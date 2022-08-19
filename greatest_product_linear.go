package main

import (
	"fmt"
	"math"
)

// pattern: highest two or lowest two
func highestProduct(arr []float64) float64 {
	highest := math.Inf(-1)
	high := math.Inf(-1)
	lowest := math.Inf(1)
	low := math.Inf(-1)

	for _, num := range arr {
		if num >= highest {
			high = highest
			highest = num
		} else if num >= high {
			high = num
		}

		if num <= lowest {
			low = lowest
			lowest = num
		} else if num <= low {
			low = num
		}
	}

	if lowest*low > highest*high {
		return lowest * low
	} else {
		return highest * high
	}
}

func main() {
	arr := []float64{5, -10, -6, 9, 4}
	fmt.Println(highestProduct(arr)) // 60
}
