""" 121. Best Time to Buy and Sell Stock
Array for which the ith element is the price of a given stock on day i. If you were only 
  permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), 
  design an algorithm to find the maximum profit.
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Faster than 94% of Python submissions
"""

from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) < 1:
            return 0
        maxP = 0
        pBuy = prices[0]
        for p in prices:
            if p < pBuy:
                pBuy = p
            else:
                maxP = max(p - pBuy, maxP)
        return maxP
    
s = Solution()
print(s.maxProfit([7,1,5,3,6,4])) # 5
print(s.maxProfit([7,6,4,3,1])) # 0
    
