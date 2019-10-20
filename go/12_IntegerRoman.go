/* 12. Integer to Roman
  Faster than 73% of Go submissions
*/

package main

import "fmt"

func main() {
	fmt.Println(intToRoman(4)) // "IV"
	fmt.Println(intToRoman(9)) // "IX"
	fmt.Println(intToRoman(58)) // "LVIII"
	fmt.Println(intToRoman(1994)) // "MCMXCIV"
}

var ones = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var hund = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var thou = []string{"", "M", "MM", "MMM", "MMMM"}
var mapping = [][]string{ones, tens, hund, thou}
func intToRoman(num int) string {
	out := ""
	for i:=0; num!=0; i++ {
		out = mapping[i][num%10] + out
		num /= 10
	}
	return out
}

