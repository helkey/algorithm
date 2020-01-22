/* 371. Sum of Two Integers
Calculate sum of two integers a and b, not using operator + and -.
https://leetcode.com/problems/sum-of-two-integers/

Faster than 100% of Go submissions
*/
 
package main
import "fmt"

func main() {
	// fmt.Println(getSum(-1, -3)) // -4
	// fmt.Println(getSum(1, 15)) // 16
	fmt.Println(getSum(-1, 1)) // 0

}


// Repeatedly generate bitwise sum plus all carry bits,
//  repeat until carry bits are all zero
func getSum(a int, b int) int {
	nbits := 64
	for i:=0; (a != 0) && (i < nbits); i++ {
		a, b = bit_sum(a, b)
		fmt.Println(a, b)
	}
	return b
}

// Bitwise sum plus shifted carry bits
func bit_sum(a int, b int) (int, int) {
	mask := -1 // 1<<nbits - 1
	bit_sum := a ^ b // sum given by xor if no carries between bits
	bit_carry := ((a & b) << 1) & mask // carry bits
	return bit_carry, bit_sum
}
	
