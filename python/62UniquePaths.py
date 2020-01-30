""" 62. Unique Paths
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?
  https://leetcode.com/problems/unique-paths/

Here this can be calculated by the number of permutations, i.e.:
  'rrdddr', 'rdddrr',...  =
  (m + n - 2)!/((m-1)! * (n-1)!) https://brilliant.org/wiki/permutations-with-repetition/

Faster than 70% of Python submissions
"""

from math import factorial

class Solution:
    def uniquePaths(self, m: int, n: int) -> int:     
        if m <= 1 or n <= 1:
            return 1
        fM = factorial(m - 1)
        fN = fM if n == m else factorial(n - 1)
        unique = int(factorial(m + n - 2) / (fM * fN))
        return unique

s = Solution()
print(s.uniquePaths(7, 1)) # assert = 1
print(s.uniquePaths(3, 2)) # assert = 3
print(s.uniquePaths(7, 3)) # assert = 28


        
