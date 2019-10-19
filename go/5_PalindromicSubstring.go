/* 5. Longest Palindromic Substring
(faster than 95% of Go submissions)
*/

package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxL := 1
	sStart := 0
	for i:=1; i<len(s); i++ {
		// Search odd palindrome
		// (could combine common code for even/odd palindromes)
		for j:=i-1; j>=0; j-- {
			k := 2 * i - j
			if (k >= len(s)) || (s[j] != s[k]) {
				break
			}
			if k - j + 1 > maxL {
				maxL = k - j + 1
				sStart = j
			}
			
		}
		// Search even palindrome
		for j:=i-1; j>=0; j-- { 
			k := 2 * i - j - 1
			if (k >= len(s)) || (s[j] != s[k]) {
				break
			}
			if k - j + 1 > maxL {
				maxL = k - j + 1
				sStart = j
			}
		}
	}
	return s[sStart:sStart+maxL]
}

