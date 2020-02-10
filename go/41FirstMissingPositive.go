/* 41. First Missing Positive
Given an unsorted integer array, find the smallest missing positive integer.
  https://leetcode.com/problems/first-missing-positive/

Faster than 100% of Go submissions
*/

package main
import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1,2,0})) // assert 3
	fmt.Println(firstMissingPositive([]int{3,4,-1,1})) // assert 2
	fmt.Println(firstMissingPositive([]int{7,8,9,11,12})) // assert 1
	fmt.Println(firstMissingPositive([]int{1,2,3})) // assert 4
	fmt.Println(firstMissingPositive([]int{})) // assert 1
}

func firstMissingPositive(nums []int) int {
	// Remove numbers to be ignored
	for i, v := range(nums) {
		if v <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	// Indicate number found using sign bit/array position
	for _, v := range(nums) {
		indx := abs(v) - 1
		if indx < len(nums) && nums[indx] > 0 {
			nums[indx] = - nums[indx]
		}
	}
	// Missing number is position of first non-negative integer
	for i, v := range(nums) {
		if v > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
