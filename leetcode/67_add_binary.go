// 67_add_binary
// Faster than 100.00% of Go online submissions for Add Binary.

package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))      // 100
	fmt.Println(addBinary("1010", "1011")) // 10101
	fmt.Println(addBinary("110", "0"))     // 110
	fmt.Println(addBinary("0", "0"))       // 0
}

func addBinary(a string, b string) string {
	len_a := len(a)
	len_b := len(b)
	lenCalc := max(len_a, len_b) + 1
	sumArr := make([]byte, lenCalc)
	var bit_a, bit_b, carry byte
	// Calculate lsb to msb
	for i := 1; i <= lenCalc; i++ {
		bit_a = 0
		if i <= len_a {
			bit_a = a[len_a - i] - '0'
		}
		bit_b = 0
		if i <= len_b {
			bit_b = b[len_b - i] - '0'
		}
		// Carry/overflow
		bit_sum := bit_a + bit_b + carry
		carry = 0
		if bit_sum > 1 {
			bit_sum -= 2
			carry = 1
		}
		// Bit reverse sum (msb first)
		sumArr[lenCalc - i] = bit_sum + '0'
	}
	// Suppress leading zeros, convert byte array to string
	iStart := 0
	for ; (sumArr[iStart] == '0') && (iStart < lenCalc - 1); iStart++ {
	}
	return string(sumArr[iStart: len(sumArr)])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
