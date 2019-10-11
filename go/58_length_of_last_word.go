// 35_...
// faster than 100% of Go online submissions

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World")) // 5
	fmt.Println(lengthOfLastWord(" World")) // 5
	fmt.Println(lengthOfLastWord("World")) // 5
	fmt.Println(lengthOfLastWord("a ")) // 1
}

func lengthOfLastWord(s string) int {
	space := byte(' ')
	lastSpace := -1
	nextSpace := -1
	lastChar := -1
	for i:=0; i<len(s); i++ {
		if s[i] == space {
			nextSpace = i
		} else {
			lastChar = i
			lastSpace = nextSpace
		}
	}
	if lastChar == -1 {
		return 0
	}
	return lastChar - lastSpace
}
