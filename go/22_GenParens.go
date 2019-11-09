/* 22. Generate Parentheses
  Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

n=3: ["((()))",
      "(()())",
      "(())()",
      "()(())",
      "()()()" 

Faster than 90% of Go submissions */

package main
      
import "fmt" 

func main() {
	// fmt.Println(generateParenthesis(2))
	out := generateParenthesis(3)
	for _, v := range out {
		fmt.Println(v)
	}}

// Recursive solution
func generateParenthesis(n int) []string {
	strs := []string{}
	fAdd := func(str string) {
		strs = append(strs, str)
	}
	if n < 1 {
		return []string{}
	}
	// start with n left parens, zero right parens
	genParen(n, 0, "", fAdd)
	return strs
}


/* Two paths:
    Start opening bracket if one (of n) left to place
Start w/closing bracket if not exceeding number of opening brackets */
func genParen(nLft int, nRgt int, str string, fAdd func(string)) {
	// fmt.Println(nLft, nRgt, str)
	if nLft == 0 && nRgt == 0 {
		fAdd(str)
	}
	if nLft > 0 {
		// fmt.Println("nLft-1")
		genParen(nLft - 1, nRgt + 1, str + "(", fAdd)
	}
	if nRgt > 0 {
		genParen(nLft, nRgt - 1, str + ")", fAdd)
	}
	return
}
