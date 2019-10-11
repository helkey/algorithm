/** 53. Maximum Subarray

Integer array nums, find the contiguous subarray (containing at least one number)
  which has the largest sum

O(n) solution; divide and conquer solution

**/

package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // ans=6
}

/// DYNAMIC PROGRAMMING / MEMOZIATION O(n); Space complexity O(n)
// Make array of sums of length n; use to compute sums of length n+1
func maxArrayN(nums []int, sumA []int, n int) []int {
	nLen := len(sumA) - 1
	newSum := make([]int, nLen)
	for i := 0; i < nLen; i++ {
		newSum[i] = sumA[i] + nums[i+n]
	}
	fmt.Println(newSum)
	return newSum
}

/** Iterative solution; O(n^2) Space complexity O(1)
  SLOW! Runtime: 248 ms, faster than 1.00% of Go online submissions for Maximum Subarray.
        Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Maximum Subarray **/
func maxSubArray(nums []int) int {
	stMax := -1
	var max int
	// var endMax int
	for st := 0; st < len(nums); st++ {
		sum := 0
		for end := st; end < len(nums); end++ {
			sum += nums[end]
			if (stMax < 0) || (sum > max) {
				max = sum
				stMax = st
				// end_max = end
			}
		}
	}
	return max
}
