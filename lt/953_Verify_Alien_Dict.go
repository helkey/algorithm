package main

import (
	"fmt"
)

func main() {
	testSorted()
}

func isAlienSorted(words []string, order string) bool {
	position := find_position(order)
	for i := 0; i < len(words)-1; i++ {
		if !isAlphabetical(words[i], words[i+1], position) {
			return false
		}
	}
	return true
}

const numChar = 26

// Position in alphabet for each character
// NOTE: Would use hash table if alphabet uses runes
func find_position(order string) (position [numChar]int) {
	// position = make([]int, numChar)
	for indx, b := range order {
		position[b-'a'] = indx
	}
	// fmt.Println(position)
	return
}

func isAlphabetical(word1 string, word2 string, position [numChar]int) bool {
	lenMin := min(len(word1), len(word2))
	for i := 0; i < lenMin; i++ {
		lett1 := position[word1[i]-'a']
		lett2 := position[word2[i]-'a']
		if lett1 == lett2 {
			continue
		}
		// Letters are different
		return lett1 < lett2
	}
	// Common stem from both words are equal
	// If alphabetical, word2 should be longer
	return len(word1) < len(word2)
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

func testSorted() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order), words, order)

	words = []string{"word", "world", "row"}
	order = "worldabcefghijkmnpqstuvxyz"
	fmt.Println(isAlienSorted(words, order), words, order)

	words = []string{"apple", "app"}
	order = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order), words, order)

	words = []string{"hello", "leetcode"}
	order = "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order), words, order)
}
