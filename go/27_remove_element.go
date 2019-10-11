// 27_remove_element.go
// Faster than 100.00% of Go online submissions for Remove Element.

package main

import "fmt"

func main() {
	// fmt.Println(removeElement([]int{3,2,2,3}, 3)) // len=2
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2)) // len=5
}

func removeElement(nums []int, val int) int {
	j := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	// fmt.Println(nums)
	return j
}
