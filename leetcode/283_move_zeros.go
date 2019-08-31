

package main

import "fmt"


func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums) // [1,3,12,0,0]
}
	
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
