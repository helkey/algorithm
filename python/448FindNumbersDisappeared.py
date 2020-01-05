""" 448. Find All Numbers Disappeared in an Array
  Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and 
others appear once. Find all the elements of [1, n] inclusive that do not appear in this array.
 https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

HARD!: Do it without extra space and in O(n) runtime. Assume returned list does not count as extra space.
  Faster than 43% of Python submissions
  (NOTE: faster submissions are using extra space for a counting array)
"""

from typing import List

class Solution:
        def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
                for i, num in enumerate(nums):
                        num = abs(num)
                        val = nums[num-1]
                        # track by setting number negative
                        if val > 0:
                                nums[num-1] = -val
                return [i+1 for i, num in enumerate(nums) if num > 0]
        
s = Solution()
print(s.findDisappearedNumbers([4,3,2,7,8,2,3,1])) # [5, 6]
