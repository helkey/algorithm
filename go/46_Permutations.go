/* 46. Permutations
   Given a collection of distinct integers, return all possible permutations.

Faster than 93% of Go submissions
*/

package main
import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2}))
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{5,4,6,2})) // 
}

// Backtracking solution
func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	perms := make([][]int, 0)
	add_perm := func(perm []int)() {
		permCpy := make([]int, len(nums))
		copy(permCpy, perm) // need to explicitly copy
		perms = append(perms, permCpy)
		return
	}
	gen_perm(nums, 0, add_perm)
	return perms
}

// Backtracking iteration of generating a permuation of nums
func gen_perm(nums []int, i int, add_perm func([]int)) {
	if i == len(nums) {
		add_perm(nums)
		return
	}
	gen_perm(nums, i+1, add_perm)
	for j := i + 1; j < len(nums); j++ {
		nums[i], nums[j] = nums[j], nums[i]
		gen_perm(nums, i+1, add_perm)
		nums[i], nums[j] = nums[j], nums[i]
	}		
}		

