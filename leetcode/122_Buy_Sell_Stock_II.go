/** 122. Best Time to Buy and Sell Stock II

Multiple transactions (one/day)

Runtime: 4 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock II.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock II.
*/

package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4})) // 7
	fmt.Println(maxProfit([]int{1,2,3,4,5})) // 4
	fmt.Println(maxProfit([]int{7,6,4,3,1})) // 0
	fmt.Println(maxProfit([]int{0, 1})) // 1
}

func maxProfit(prices []int) int {
	profit := 0
	for s:=0; s<len(prices)-1; s++ {
		delta := prices[s+1] - prices[s]
		if delta > 0 {
			profit += delta
		}
	}
	return profit
}


