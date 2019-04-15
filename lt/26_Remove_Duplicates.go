/** 26. Remove Duplicates from Sorted Array

Given a sorted array nums, remove the duplicates in-place 
such that each element appear only once and return the new length.

Runtime: 60 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 7.5 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.

**/

package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})) // 5
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3})) // 4
	fmt.Println(removeDuplicates([]int{0})) // 1
	fmt.Println(removeDuplicates([]int{})) // 0
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	n := 1
	for _, v := range nums {
		if v != nums[n-1] {
			nums[n] = v
			n++
		}
	}
	return n
		
}
