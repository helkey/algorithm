/** 13. Roman to Integer

**/

package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III")) // 3
	fmt.Println(romanToInt("IV")) // 4
	fmt.Println(romanToInt("IX")) // 9
	fmt.Println(romanToInt("LVIII")) // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}


func romanToInt(s string) int {
	sum := 0
	code := map[byte]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}
	for i:=0; i<len(s); i++ {
		val := code[s[i]]
		end := (i == len(s) - 1)
		if !end && (code[s[i+1]] > val) {
			val = -val
		}
		sum += val
	}
	return sum
}
