/* 268. Missing Number
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Faster than 35% of Go submissions. XOR approach might be faster.
  (Setting up additional array and looking for missing number likely faster, but not O(1) space complexity)
*/

package main
import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1})) //  2
	fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1})) //  8
}

// Algorithm should run in linear runtime complexity... implement it using only constant extra space complexity
func missingNumber(nums []int) int {
	// Gauss formula: Sum [0.. n] = n * (n-1) / 2
	n := len(nums) + 1
	sum := n * (n - 1) / 2
	for _, i := range nums {
		sum = sum - i
	}
	return sum // remainder is the missing number
}

	
	
