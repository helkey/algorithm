/** 88. Merge Sorted Array

e.g.
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
Output: [1,2,2,3,5,6]

**NOTE: may be faster to do multiple indexing into array, vs defining a new variable
  this solution has improvement over published solutions by stopping when n==0
Runtime: 4 ms, faster than 21.60% of Go online submissions for Merge Sorted Array.
Memory Usage: 3.6 MB, less than 100.00% of Go online submissions for Merge Sorted Array.
**/

package main

import "fmt"

func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	merge([]int{0}, 0, []int{1}, 1)
	merge([]int{1}, 1, []int{}, 0)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	// sort from back to front
	m--
	n--
	for p:= len(nums1)-1; n>=0; p-- {
		if n<0 {
			n1 := nums1[m]
			nums1[p] = n1
			m--
		} else if m<0 {
			n2 := nums2[n]
			nums1[p] = n2
			n--
		} else {
			n1 := nums1[m]
			n2 := nums2[n]
			if n1 > n2 {
				nums1[p] = n1
				m--
			} else {
				nums1[p] = n2
				n--
			}
		}
		
	}
	fmt.Println(nums1)
}
