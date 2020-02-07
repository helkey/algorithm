/* 55. Jump Game
Given an array of non-negative integers, you are initially positioned at the first index of the array.
  Each element in the array represents your maximum jump length at that position.
  Determine if you are able to reach the last index.
https://leetcode.com/problems/jump-game/

Faster than 96% of Go submissions (O(N))
*/

package main
import "fmt"

func main() {
	fmt.Println(canJump([]int{2,3,1,1,4})) // assert true
	fmt.Println(canJump([]int{3,2,1,0,4})) // assert false
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	maxRange := 0
	// Early termination if maxRange large enough
	for i:=0; (i < len(nums) && maxRange < len(nums) - 1); i++ {
		if i > maxRange {
			return false
		}
		maxRange = max(maxRange, i + nums[i])
		// fmt.Println(i, nums[i], maxRange)
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
