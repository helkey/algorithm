// Backtracking solutions: https://www.youtube.com/watch?v=78t_yHuGg-0
// https://leetcode.com/problems/permutations/discuss/18239/a-general-approach-to-backtracking-questions-in-java-subsets-permutations-combination-sum-palindrome-partioning

/* 40. Combination Sum II
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

*/

package main
import (
	"fmt"
	"sort"
)

func main() {
	// lst = 
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
}

// Faster than 10% of Go online submissions
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	combos := [][]int{}
	var vLast int
   	for i, v := range(candidates) {
		if (i != 0) && (vLast == v) {
			continue
		}
		vLast = v
		if (v < target) && (i < len(candidates)) {
			// Don't use candidates < v, or will get repeated solutions
			candidates_r := candidates[i + 1: len(candidates)] // remaining
			for _, sublist := range combinationSum2(candidates_r, target - v) {
				combos = append(combos, append([]int{v}, sublist...))
			}
		} else if v == target {
				combos = append(combos, []int{v})
		}
	}
	return combos
}
