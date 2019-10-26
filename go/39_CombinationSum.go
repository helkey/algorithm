/* 39. Combination Sum
  Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), 
   find all unique combinations in candidates where the candidate numbers sums to target.

(faster than 86% of Go submissions) */

package main
import "fmt"

func main() {
	// lst = 
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
	fmt.Println(combinationSum([]int{2,3,5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	combos := [][]int{}
	for i, v := range(candidates) {
		if v < target {
			// Don't use candidates < v, or will get repeated solutions
			candidates_r := candidates[i: len(candidates)] // remaining
			for _, sublist := range combinationSum(candidates_r, target - v) {
				combos = append(combos, append([]int{v}, sublist...))
			}
		} else if v == target {
				combos = append(combos, []int{v})
		}
	}
	return combos
}
