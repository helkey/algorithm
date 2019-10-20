/* 17. Letter Combinations of a Phone Number
   Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

Faster than 100% of Go submissions */

package main

import "fmt"

func main() {
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("23"))
}

var num2char = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
func letterCombinations(digits string) []string {
	if len(digits) == 0 || digits[0] <= 1 {
		return []string{}
	}
	var output []string
	chars := num2char[byte(digits[0]) - '0']
	for i:=0; i<len(chars); i++ {
		substrs := letterCombinations(digits[1:len(digits)])
		if len(substrs) == 0 {
			output = append(output, string(chars[i]))
		} else {
			for _, substr := range substrs {
				output = append(output, string(chars[i]) + substr)
			}
		}
	}
	return output
}

