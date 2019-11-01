/* 11. Container With Most Water
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). 
  n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). 
  Find two lines, which together with x-axis forms a container, such that container contains the most water.

(faster than 90% of Go submissions) */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

// Eliminate smallest outer line
func maxArea(height []int) int {
	maxA := 0
	for ; len(height)>0 ; {
		// Area defined by two outer lines
		h1, hn := height[0], height[len(height) - 1]
		area := min(h1, hn) * (len(height) - 1)
		maxA = max(area, maxA)
		if h1 < hn {
			height = height[1:]
		} else {
			height = height[:len(height)-1]
		}
	}
	return maxA
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}


func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}


/* N^2 solution
func max_N2(height []int) int {
	maxWater := 0
	for i, h1 := range height {
		for j, h2 := range height {
			box := abs(i - j) * min(h1, h2)
			maxWater = max(maxWater, box)
		}
	}
	return maxWater
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
} */
