/* 387. First Unique Character in a String
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
 https://leetcode.com/problems/first-unique-character-in-a-string/

Fasterthan 95% of Go submissions
*/


package main
import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode")) // 0
	fmt.Println(firstUniqChar("loveleetcode")) // 2
	fmt.Println(firstUniqChar("eee")) // -1
}

func firstUniqChar(s string) int {
	char := [256]int{}
	for i:=0; i<len(s); i++ {
		if char[s[i]] < 2 {
			char[s[i]] += 1
		}
	}
	for i:=0; i<len(s); i++ {
		if char[s[i]] == 1 {
			return i
		}
	}
	err_val := -1
	return err_val
}
