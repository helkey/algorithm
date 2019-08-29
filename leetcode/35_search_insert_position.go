package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5)) // 2
}

func searchInsert(nums []int, target int) int {
    nNums := len(nums)
	for i:=0; i<nNums; i++ {
		if nums[i] >= target {
			return i
		}
	}
	return nNums
}

