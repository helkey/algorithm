# 217. Contains Duplicate
# Faster than 82% of Python3 submissions

from typing import List


class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        prev = {}
        for n in nums:
            if n in prev:
                return True
            prev[n] = True
        return False

s = Solution()
print(s.containsDuplicate([1,2,3,1])) # True
print(s.containsDuplicate([1,2,3,4])) # False
