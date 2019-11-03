/* 75. Sort Colors
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent,
  with the colors in the order red, white and blue.

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.

Q: Could you come up with a one-pass algorithm using only constant space?
  {A: No, providedd solution is worse than this 2-pass counting solution}

(Faster than 100% of Go submissions) */

package main

import "fmt"

func main() {
	in := []int{2, 0, 2, 1, 1, 0}
	sortColors(in)
	fmt.Println(in) // [0,0,1,1,2,2]
}

// O(n) Two-pass; One pass through, count number of each occurance
//  Second pass; write out same number of elements as in first pass
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	count := make(map[int]int)
	iMin := 0
	iMax := 0
	for indx, num := range nums {
		if _, ok := count[num]; !ok {
			count[num] = 0
		}
		count[num] += 1
		if indx == 0 {
			iMin, iMax = num, num
		} else {
			iMin = min(iMin, num)
			iMax = max(iMax, num)
		}
	}
	indx := 0
	for iColor := iMin; iColor <= iMax; iColor++ {
		for j := 0; j < count[iColor]; j++ {
			nums[indx] = iColor
			indx += 1
		}
	}
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
