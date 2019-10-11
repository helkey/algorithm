/** 121. Best Time to Buy and Sell Stock

One transaction (i.e., buy one and sell one share of the stock), 
design an algorithm to find the maximum profit.

DO BETTER: contruct array of local min & local max prices

Runtime: 4 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Best Time to Buy and Sell Stock.
*/

package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4})) // 5
	fmt.Println(maxProfit([]int{7,6,4,3,1})) // 0
	fmt.Println(maxProfit([]int{0, 1})) // 1
}

func maxProfit(prices []int) int {
	if (len(prices) <= 1) || decreasing(prices) {
		return 0
	}
	maxPr := 0
	for s:=0; s<len(prices)-1; s++ {
		if prices[s] < prices[s+1] {
			for e:=s+1; e<len(prices); e++ {
				profit := prices[e] - prices[s]
				if profit > maxPr {
					maxPr = profit
				}
			}
		}
	}
	return maxPr
}

func decreasing(p []int) bool {
	for i:=1; i<len(p); i++ {
		if p[i] > p[i-1] {
			return false
		}
	}
	return true
}

