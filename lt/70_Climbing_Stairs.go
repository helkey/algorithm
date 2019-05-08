/** 70. Climbing Stairs

You are climbing a stair case. 
It takes n steps to reach to the top.
Each time you can either climb 1 or 2 steps. 
In how many distinct ways can you climb to the top?

Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Climbing Stairs.
**/

package main

import "fmt"

func main() {
	for n:=1; n<5; n++ {
		fmt.Println(n, climbStairs(n))
	}
}

func climbStairs(n int) int {
	ways := []int{1, 1}
	for i:=2; i<=n; i++ {
		lenW := len(ways)
		ways = append(ways, ways[lenW-2] + ways[lenW-1])
	}
	return ways[len(ways)-1]
}
