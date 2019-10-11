/* 35. Search Insert Position
   Given a sorted array and a target value, return the index if the target is found.
  If not, return the index where it would be if it were inserted in order */

package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2
}

// Use Binary search! O(ln(N))
func searchInsert(nums []int, target int) int {
	nSt, nE := 0, len(nums)
	for ; nSt >= nE - 1; nMid = (nSt + nE)/2 {
		if nums[nMid] == target {
			return nMid
		} else {
			nSt = nMid
		} else {
			nE = nMid
		}
	}
	// final check here return nNums
}

// not linear search O(N)
func searchInsert(nums []int, target int) int {
	nNums := len(nums)
	for i := 0; i < nNums; i++ {
		if nums[i] >= target {
			return i
		}
	}
	return nNums
}
