package main

import (
	"fmt"
)

func groupArray(arr []string) []string {
	// convert the array of strings into a map (k-v: str-count)
	// iterate through the map
	// 	 - push each key v times in a new arr
	// return the new arr

	hash := map[string]int{}
	newArr := []string{}

	for _, str := range arr {
		_, ok := hash[str]
		if ok {
			hash[str] += 1
		} else {
			hash[str] = 1
		}
	}

	for k, v := range hash {
		for i := 0; i < v; i++ {
			newArr = append(newArr, k)
		}
	}

	return newArr
}

func main() {
	arr := []string{"a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d"}
	fmt.Println(groupArray(arr))
}
