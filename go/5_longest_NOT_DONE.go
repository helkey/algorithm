//DYNAMIC PROGRAMMING (NOT CLEAR WHY)

// 5. Longest Palindromic Substring
// Check from center outward, stop when edges limit finding longer palendrome

package main

import (
	"fmt"
	"strconv"
)

func main() {
}

// Copied from 9._***
func isPalindrome(x int) bool {
	j := len(xStr) - 1
	for i := 0; i < (len(xStr) / 2); i++ {
		if xStr[i] != xStr[j] {
			return false
		}
		j--
	}
	return true
}
