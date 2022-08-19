package main

import (
	"fmt"
	"reflect"
)

func areAnagrams(s1, s2 string) bool {
	// convert each string to a hash (chr, count areAnagrams)
	// compare two hashes; return true if fully matched

	hash1 := map[rune]int{}
	hash2 := map[rune]int{}

	for _, chr := range s1 {
		_, ok := hash1[chr]
		if ok {
			hash1[chr] += 1
		} else {
			hash1[chr] = 1
		}
	}

	for _, chr := range s2 {
		_, ok := hash2[chr]
		if ok {
			hash2[chr] += 1
		} else {
			hash2[chr] = 1
		}
	}

	return reflect.DeepEqual(hash1, hash2)
}

func main() {
	str1 := "hello"
	str2 := "olehl"
	str3 := "hihih"
	fmt.Println(areAnagrams(str1, str2)) // true
	fmt.Println(areAnagrams(str1, str3)) // false
}
