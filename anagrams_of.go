package main

import "fmt"

/*
find all anagrams
- given a string, return an array of anagrams of the given string (inclusive)
- let's say we have a working anagramOf function that has found 3-letter anagrams for a 4-letter word
  - the 4th letter will be placed at the every gap (before/after/inbetween characters)

helper
- insert the chr to the word at all gaps
  - prepend the chr to the word and push to the result arr
  - iterate through the word
    - insert the chr after the current idx
      - push to the result arr
  - return the result
*/

func insertChrToWord(word, chr string) []string {
	result := []string{}
	for i := 0; i < len(word); i += 1 {
		result = append(result, word[0:i]+chr+word[i:])
	}
	result = append(result, word+chr)
	return result
}

func anagramsOf(word string) []string {
	result := []string{}
	if len(word) == 1 {
		result = append(result, word)
		return result
	}

	substrings := []string{}
	substrings = anagramsOf(word[1:])
	for _, substr := range substrings {
		result = append(result, insertChrToWord(substr, string(word[0]))...)
	}
	return result
}

func main() {
	fmt.Println(anagramsOf("abc")) // ["abc", "acb", "bac", "bca", "cab", "cba"]
}
