/** 20. Valid Parentheses

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
  *Open brackets must be closed by the same type of brackets.
  *Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

**/

package main

import "fmt"

func main() {
	fmt.Println(isValid("()")) // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]")) // false
	fmt.Println(isValid("([)]")) // false
	fmt.Println(isValid("{[]}")) // true
	fmt.Println(isValid("")) // true
	
}

func isValid(s string) bool {
	start := map[byte]bool{'(':true, '{':true, '[':true}
	// stop :=  map[byte]bool{')':true, '}':true, ']':true}
	match := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	open := []byte{}
	for i:=0; i<len(s); i++ {
		si := s[i]
		if start[si] {
			open = append(open, si) // new open character
			// fmt.Println("new", open, si)
		} else {
			lenOpen := len(open)
			if (lenOpen > 0) && (si == match[open[lenOpen-1]]) {
				// matching closing character
				open = open[:lenOpen-1]
				// fmt.Println("matching", open, string(si))
			} else {
				// fmt.Println("not matching", string(si))
				// fmt.Println(open[lenOpen-1], match[si])
				return false // not matching
			}
		}
	}
	// check if all open characters matched with closing characters
	// fmt.Println("check if all chars matched")
	return (len(open) == 0) 
}

