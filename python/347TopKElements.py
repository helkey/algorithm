""" 347. Top K Frequent Elements
Given a non-empty array of integers, return the k most frequent elements.
Algorithm's time complexity must be better than O(n log n), where n is the array's size.

Heapify: O(n)
k-smallest: O(k * log(n))

Faster than 83% of Python submissions
"""

from collections import defaultdict
import heapq
from typing import List

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        ndict = defaultdict(int)
        for n in nums:
            ndict[n] += 1
        # Negate frequency to use min-heap implementation
        ncount = [[-ndict[k], k] for k in ndict]
        
        # ncount.sort()
        # return [k for _, k in ncount[:k]]
        return [k for _, k in heapq.nsmallest(k, ncount)]


s = Solution()
print(s.topKFrequent([1,1,1,2,2,3], 2))
print(s.topKFrequent([3,1,1,1,2,2], 2))
