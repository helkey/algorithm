/* Given a list of non-negative integers representing the amount of money of each house,
   determine the maximum amount of money you can rob without 2 adjacent houses */

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums), 4) // 4
	nums = []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums), 12) // 12
}

func rob(nums []int) int {
	nNums := len(nums)
	if nNums == 0 {
		return 0
	}
	if nNums == 1 {
		return nums[0]
	}

	// Generate max sum (in array 'nums')
	if len(nums) > 2 {
		nums[2] += nums[0]
	}
	for i := 3; i < nNums; i++ {
		nums[i] += max(nums[i-3], nums[i-2])
	}

	// Return maximum value (one of last 2 values)
	return max(nums[nNums-2], nums[nNums-1])
}

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}
