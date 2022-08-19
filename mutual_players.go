package main

import "fmt"

func mutualPlayers(basketball, football []map[string]string) []string {
	// convert both maps into maps with full name
	// if matched, append to a new array
	// return the new array
	hash1 := map[string]bool{}
	hash2 := map[string]bool{}
	result := []string{}

	for _, hash := range basketball {
		hash1[hash["first_name"]+" "+hash["last_name"]] = true
	}

	for _, hash := range football {
		hash2[hash["first_name"]+" "+hash["last_name"]] = true
	}

	for key := range hash1 {
		if hash2[key] {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	basketball_players := []map[string]string{
		{"first_name": "Jill", "last_name": "Huang", "team": "Gators"},
		{"first_name": "Janko", "last_name": "Barton", "team": "Sharks"},
		{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Sharks"},
		{"first_name": "Jill", "last_name": "Moloney", "team": "Gators"},
		{"first_name": "Luuk", "last_name": "Watkins", "team": "Gators"},
	}
	football_players := []map[string]string{
		{"first_name": "Hanzla", "last_name": "Radosti", "team": "32ers"},
		{"first_name": "Tina", "last_name": "Watkins", "team": "Barleycorns"},
		{"first_name": "Alex", "last_name": "Patel", "team": "32ers"},
		{"first_name": "Jill", "last_name": "Huang", "team": "Barleycorns"},
		{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Barleycorns"},
	}

	fmt.Println(mutualPlayers(basketball_players, football_players))
}
