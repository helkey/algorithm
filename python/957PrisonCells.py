""" 957. Prison Cells After N Days  https://leetcode.com/problems/prison-cells-after-n-days/
8 prison cells in a row, and each cell is either occupied or vacant.
  Each day, whether the cell is occupied or vacant changes according to the following rules:
  If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell 
  becomes occupied. Otherwise, it becomes vacant.
NOTE: cell0/N-1 are always zero after the first day (not clear from problem description);
  different than constraints in test.

Faster than 97% of submissions
"""

from typing import List

class Solution:   
    def prisonAfterNDays(self, cells: List[int], N: int) -> List[int]:        
        def cell_i(cells, i):
            if i == 0 or i == len(cells) - 1:
                return 0
            if cells[i-1] == cells[i+1]:
                return 1
            return 0

        nRep = 14 # repetition rate of cycles (for 8 cells)
        nCalc = N  - nRep * int((N - 1)/nRep) # number of days to get equivalent pattern
        print("nCalc", nCalc)
        for _ in range(nCalc):
            cells = [cell_i(cells, i) for i in range(len(cells))]
        return cells


s = Solution()
for i in range(7):
    pass # print(s.prisonAfterNDays([0,1,0,1,1,0,0,1], i + 1))
""" Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
    Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
    Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
    Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
    Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
    Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
    Day 7: [0, 0, 1, 1, 0, 0, 0, 0] """
print()

for i in range(13):
    pass # print(s.prisonAfterNDays([0,1,0,1,0,1,0,0], i + 1)) # [0,1,0,0,0,1,0,0]

print(s.prisonAfterNDays([1,0,0,1,0,0,0,1], 826)) # [0,1,1,0,1,1,1,0]
