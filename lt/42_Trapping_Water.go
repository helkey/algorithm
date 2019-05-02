//DON'T NEED TWO PASSES; DONT ALLOCATE 2nd ARRAY!!!
/** 42. Trapping Rain Water

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

**/

package main

import (
	"fmt"
	// "math"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // anw=6
	fmt.Println(trap([]int{0})) // anw=0
}

func trap(height []int) int {
	// Calculate water level
	t_level := make([]int, len(height))
	lo:= 0
	hi:= len(height)-1
	if hi < 0 {
		return 0	// NOTE: one of test cases has len(height)=0
	}
	level := min(height[lo], height[hi])
	t_level[lo] = level
	t_level[hi] = level
	for hi - lo > 1 {
		if (height[lo] < height[hi]) || (height[lo] < level) {
			lo++
		} else {
			hi--
		}
		level = max(level, min(height[lo], height[hi]))
		t_level[lo] = level
		t_level[hi] = level
		// fmt.Println(lo, hi, level)
	}

	// Calculate stored water
	stored := 0
	for i, h := range height {
		if t_level[i] > h {
			stored += (t_level[i] - h)
		}
	}
	return stored
}


func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
