package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121)) // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10)) // false
	fmt.Println(isPalindrome(0)) // true
	fmt.Println(isPalindrome(1000021)) //false
	
}

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	j := len(xStr)- 1
	for i:=0; i<(len(xStr)/2); i++ {
		if xStr[i] != xStr[j] {
			return false
		}
		j--
	}
	return true
}
