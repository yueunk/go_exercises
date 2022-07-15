package main

import "fmt"

// return the first non-duplicated character in a string
// ex. "minimum" -- "n" and "u" are singles; return the first single character "n" at O(N)
// convert the string to map

func firstSingleCharacter(str string) string {
	hash := map[rune]int{}

	for _, chr := range str {
		hash[chr] += 1
	}

	for _, chr := range str {
		if hash[chr] == 1 {
			return string([]rune{chr})
		}
	}

	return ""
}

func main() {
	fmt.Println(firstSingleCharacter("minimum")) // "n"
}
