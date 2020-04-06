/*283. Move Zeroes
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
 *You must do this in-place without making a copy of the array.
 *Minimize the total number of operations.
  https://leetcode.com/problems/move-zeroes/
*/


package main

import "fmt"


func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums) // [1,3,12,0,0]
}

// Move all non-zero numbers to front, fill remaining slots with '0' O(N)
func moveZeroes(nums []int)  {
	var i, iZ int
	// Move non-zero elements
	for ; i<len(nums); i++ {
		if nums[i] != 0 {
			nums[iZ] = nums[i]
			iZ++
		}
	}
	// Zero out remaining elements
	for ; iZ<len(nums); iZ++ {
		nums[iZ]=0
	}
	return
}
