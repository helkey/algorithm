/** 66. Plus One

Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Plus One.
**/

package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1,2,3})) // {1, 2, 4}
	fmt.Println(plusOne([]int{9, 9, 9})) // {1, 0, 0, 0}
	fmt.Println(plusOne([]int{0})) // {1, 0, 0, 0}
}

func plusOne(digits []int) []int {
	carry := 1
	for i:=len(digits)-1; i>=0; i-- {
		d := digits[i] + carry
		// fmt.Println(i, digits[i], carry, d)
		if d != 10 {
			digits[i] = d
			return digits
		} else {
			digits[i] = 0
			carry = 1
		}
	}
	return append([]int{1}, digits...)
}

	
