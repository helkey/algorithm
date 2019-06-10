/** 125. Valid Palindrome

Given a string, determine if it is a palindrome, 
considering only alphanumeric characters and ignoring cases.

Runtime: 4 ms, faster than 100.00% of Go online submissions for Valid Palindrome.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Valid Palindrome.
**/

package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("Abba")) // true
	fmt.Println(isPalindrome("race a car")) // false
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	end := len(s) - 1
	for st:=0; st<end; {
		if !alphaNum(s[st]) {
			st++
			continue
		} else if !alphaNum(s[end]) {
			end--
			continue
		}
		if toUpper(s[st]) != toUpper(s[end]) {
			return false
		}
		st++
		end--
	}
	return true
}

//check if alphanumeric
func alphaNum(b byte) bool {
	if ((b>='0') && (b<='9')) || ((b>='A') && (b<='Z')) || ((b>='a') && (b<='z')) {
		return true
	}
	return false
}

func toUpper(b byte) byte {
	if b>='a' {
		return b - 'a' + 'A'
	}
	return b
}
		
	
