package main

import "fmt"

// input: string; output: string
// convert the given string to map with k-v being rune-bool
// - iterate through alphabet
// 	- if the current chr doesn't exist in the map, return the val

func missingLetter(str string) string {
	hash := map[rune]bool{}
	for _, chr := range str {
		hash[chr] = true
	}

	for i := 'a'; i <= 'z'; i += 1 {
		if !hash[i] {
			s := string([]rune{i})
			return s
		}
	}
	return ""
}

func main() {
	fmt.Println(missingLetter("the quick brown box jumps over a lazy dog")) // "f"
}
