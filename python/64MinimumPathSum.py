""" 64. Minimum Path Sum
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right 
  which minimizes the sum of all numbers along its path.
Note: You can only move either down or right at any point in time.
https://leetcode.com/problems/minimum-path-sum/

Faster than 96% of Python submissions
"""

from typing import List

class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        nx, ny = len(grid[0]), len(grid)
        for y in range(ny):
            for x in range(nx):
                if x > 0 and y > 0:
                    grid[y][x] += min(grid[y-1][x], grid[y][x-1])
                elif y > 0:
                    grid[y][x] += grid[y-1][x]
                elif x > 0:
                    grid[y][x] += grid[y][x-1]
        return grid[ny-1][nx-1]


s = Solution()
grid = [[1,3,1],
        [1,5,1],
        [4,2,1]]
# grid = [[2]] # 2
print(s.minPathSum(grid)) # 7
