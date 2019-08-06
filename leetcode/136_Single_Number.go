/** 136. Single Number

Array of integers, every element appears twice except for one. 
Find that single one.

Runtime: 12 ms, faster than 100.00% of Go online submissions for Single Number.
Memory Usage: 4.8 MB, less than 100.00% of Go online submissions for Single Number.
**/

package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4,1,2,1,2})) // 4

}

// Bitwise xor function twice with same number will leave original number (parity function)
func singleNumber(nums []int) int {
	xor := 0
	for  _, n := range nums {
		xor ^= n
	}
	return xor
}


func hash(nums []int) int {
	hash := map[int]int{}
	for _, n := range nums {
		// check if found an odd number of times
		i, _ := hash[n]
		hash[n] = i + 1
	}
	for n, qnt := range hash {
		if qnt == 1 {
			return n
		}
	}
	return 0
}

