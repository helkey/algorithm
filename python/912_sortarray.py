# 912. Sort Array
# Merge Sort: Faster than 33% of Python submissions
#   Merge sort has good cache performance and [parallelizes well](https://en.wikipedia.org/wiki/Parallel_algorithm)

# Time complexity O(n log n)
# Space complexity O(n)
#   (illustration why not O(n log n): stackoverflow.com/questions/10342890/merge-sort-time-and-space-complexity)

from typing import List
class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        def merge(nums1: List, nums2: List) -> List[int]:
            """ Merge two sorted, non-empty lists """
            ln1, ln2 = len(nums1), len(nums2)
            # Allocate empty list
            nums = [None] * (ln1 + ln2)
            i1, i2 = 0, 0
            for k in range(ln1 + ln2):
                if (i2 == ln2) or ((i1 !=ln1) and (nums1[i1] < nums2[i2])):
                    nums[k] = nums1[i1]
                    i1 += 1
                else:
                    nums[k] = nums2[i2]
                    i2 += 1
            return nums
        
        if len(nums) <= 1:
            return nums
        if len(nums) == 2:
            return [min(nums), max(nums)]
        nSplit = round(len(nums)/2)
        return merge(self.sortArray(nums[:nSplit]), self.sortArray(nums[nSplit:]))

                     
            
