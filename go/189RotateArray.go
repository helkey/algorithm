/* 189. Rotate Array
Given an array, rotate the array to the right by k steps, where k is non-negative.
https://leetcode.com/problems/rotate-array/

Faster than 34% of Go submissions
  (faster to use the double-reverse method?)
*/

package main
import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6,7} 
	rotate(nums, 5)
	fmt.Println(nums) // [5,6,7,1,2,3,4]
}

func rotate(nums []int, k int)  {
	n := len(nums)
	k = ((k + n/2) % n) - n/2
	if k == 0 {
		return
	}
	if k > 0 {
		numsE := make([]int, k)
		copy(numsE, nums[n-k:n])
		copy(nums[k:n], nums[0:n-k])
		copy(nums[0:k], numsE)
	} else {
		numsS := make([]int, -k)
		copy(numsS, nums[0:-k])
		copy(nums[0:n+k], nums[-k:n])
		copy(nums[n+k:n], numsS)
	}
}


func rotate_indx(nums []int, k int)  {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	// Might be faster to check and allocate smaller half of array
	numsE := make([]int, n-k)
	copy(numsE, nums[0:n-k])
	for i:=0; i<n; i++ {
		if i < k {
			nums[i] = nums[i + n - k]
		} else {
			nums[i] = numsE[i - k]
		}
	}
}

func rotate_ptr(nPtr *[]int, k int)  {
	n := len(*nPtr)
	k = k % n
	if k == 0 {
		return
	}
	*nPtr = append((*nPtr)[k+1:n], (*nPtr)[0:k+1]...)
}
