/* 3Sum
  Given an array nums of n integers, find all unique triplets in the array which gives the sum of zero.

(faster than 55% of Go solutions) */
// https://leetcode.com/problems/3sum/discuss/404887/Python3-solution-extend-from-sum-2
// https://leetcode.com/problems/3sum/discuss/399442/Python-easy-to-understand-solution
// https://leetcode.com/problems/3sum/discuss/401650/Easy-to-understand-C%2B%2B-Solution-92ms-beats-93

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) /* [[-1,  0, 1],
                                                              [-1, -1, 2]] */
	fmt.Println(threeSum([]int{0, 0, 0}))
}

// Hashmap: take combos of 2-sums, look up negative in hashmap
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // hard to handle multiple values without sorting
	nNums := len(nums)
	vals := make(map[int]int) // hashmap of vals in nums
	// Ordered combos - don't need 1st two values in hashmap
	for i:=2; i<nNums; i++ {
		n := nums[i]
		if sum, ok := vals[n]; ok {
			vals[n] = sum + 1
		} else {
			vals[n] = 1
		}
	}

	trips := [][]int{}
	for i:=0; i<nNums-2; i++ {
		a := nums[i] // Need (a <=0) to find a set
		if (a > 0) || (i > 0 && nums[i] == nums[i-1]) {
			continue
		}
		for j:=i+1; j<nNums-1; j++ {
			if (j > i + 1) && (nums[j] == nums[j-1]) {
				continue
			}
			b := nums[j]
			if (a == -2 * b) && (nums[j+1] == b) {
				// Special case by repeated 'b' values
				// fmt.Println("special", a, b)
				trips = append(trips, []int{a, b, b})
				continue
			}
			goal := -(a + b)
			if (goal > b) {
				if _, match := vals[goal]; match {
					// fmt.Println(a, b, goal)
					trips = append(trips, []int{a, b, goal})
				}
			}
		}
	}
	return trips
}
