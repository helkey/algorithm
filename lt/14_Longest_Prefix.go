/** 14. Longest Common Prefix

**/

package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"})) // ""
	fmt.Println(longestCommonPrefix([]string{"d","da",""})) // ""
	fmt.Println(longestCommonPrefix([]string{"","a","b"})) // ""
	fmt.Println(longestCommonPrefix([]string{})) // ""
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	var b byte
	if len(strs) == 0 {
		return prefix
	}
	for i:=0;; i++ {
		for s, str := range strs {
			// fmt.Println(i, s, str)
			if i == len(str) {
				return prefix
			}
			if s == 0 {
				b = str[i]
			} else {
				if str[i] != b {
					return prefix
				}
			}
		}
		prefix = prefix + string(b)
	}
}
	
