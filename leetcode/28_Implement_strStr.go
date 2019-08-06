/** 28. Implement strStr()

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Implement strStr().
**/

package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll")) // 2
	fmt.Println(strStr("aaaaa", "bba")) // -1
	fmt.Println(strStr("a", "a")) // 0
	fmt.Println(strStr("aa", "")) // 0
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	lenNeedle := len(needle)
	lSearch := len(haystack)-lenNeedle+1
	for i:=0; i<lSearch; i++ {
		if needle == haystack[i:i+lenNeedle] {
			return i
		}
	}
	return -1
}

