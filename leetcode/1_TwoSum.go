// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution,
//   and you may not use the same element twice.

// Approach 1: Brute Force (apparently 80% Leetcode solutions used this approach)
// 2 nested loops; time O(n2); space O(1)

// Approach 2: Two-pass hash table
//  time O(n), space O(n)

// Approach 3: Line-pass hash table
// time O(n), space O(n)

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

// NOTE: excludes invalid edge case target == 2*i??
func twoSum(nums []int, target int) []int {
	match := make(map[int]int)
	for indx, i := range nums {
		pair := target - i
		if indx2, pres := match[pair]; pres {
			return []int{indx2, indx}
		}
		match[i] = indx
	}
	return []int{}
}

