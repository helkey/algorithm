/* 3. Longest Substring Without Repeating Characters
  Given a string, find the length of the longest substring without repeating characters.
  (Faster than 47% of Go submissions)
*/

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("b"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("tex"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	maxL := 0
	iStart := 0
	chars := make([]int, 256)
	for i:=0; i<len(s); i++ {
		c := s[i]
		if chars[c] == 0 {
			chars[c] = i + 1
			maxL = max(maxL, i - iStart + 1)
		} else {
			// Reset i to just after 1st repeated character
			//  (alternatively zero out any value > iStart)
			i = chars[c] - 1
			iStart = i + 1
			chars = make([]int, 256)
		}
	}
	return maxL
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}	
