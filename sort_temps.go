package main

import "fmt"

// sort temps ranging between 97.0 and 99.0
func sortTemps(temps []float64) []float64 {
	// store the temps into a hash (k-v being temp-count)
	// iterate through the possible temps
	// - push the existing temp from the hash into the new array
	hash := map[float64]int{}
	min := 97.0
	max := 99.0
	sorted := []float64{}

	for _, temp := range temps {
		_, ok := hash[temp]
		if ok {
			hash[temp]++
		} else {
			hash[temp] = 1
		}
	}

	for i := min * 10; i <= max*10; i++ {
		_, ok := hash[i/10]
		if ok {
			for j := 0; j < hash[i/10]; j++ {
				sorted = append(sorted, i/10)
			}
		}
	}

	return sorted
}

func main() {
	temps := []float64{98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1}
	fmt.Println(sortTemps(temps))
}
