""" 1057. Campus Bikes
On a campus represented as a 2D grid, there are N workers and M bikes, with N <= M. Each worker and bike is a 2D coordinate on this grid.
Our goal is to assign a bike to each worker. Among the available bikes and workers, we choose the (worker, bike) pair with the shortest Manhattan distance 
  between each other, and assign the bike to that worker. (If there are multiple (worker, bike) pairs with the same shortest Manhattan distance, 
 we choose the pair with the smallest worker index; if there are multiple ways to do that, we choose the pair with the smallest bike index). 
  We repeat this process until there are no available workers.
  https://leetcode.com/problems/campus-bikes/

The Manhattan distance between two points p1 and p2 is Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|.
Return a vector ans of length N, where ans[i] is the index (0-indexed) of the bike that the i-th worker is assigned to.

Heap solution much better - don't need to pull all the sorted values from array. Solution might be closer
  to O(N*logN) than O(N^2*logN)
(Even better? array of heaps, one per bike)

Heap solution (beats 100% in time/space): https://leetcode.com/problems/campus-bikes/discuss/473946/Python-heap-solution-explained-beats-100-in-time-and-space
"""

import heapq
from typing import List


# 0.4s, faster than 100%
class Solution:
    """ Find nearest bike for each worker, select closest match, repeat """
    def assignBikes(self, workers: List[List[int]], bikes: List[List[int]]) -> List[int]:
        bikeUsed = set()
        def closestBike(iWorker):
            xWork, yWork = workers[iWorker]
            minDist = None
            for iBike, bike in enumerate(bikes):
                if iBike not in bikeUsed:
                    xBike, yBike = bike
                    dist = abs(xWork - xBike) + abs(yWork - yBike)
                    if minDist == None or dist < minDist:
                        minDist = dist
                        minBike = iBike
            return minDist, minBike

        distHeap = []
        for iWorker, worker in enumerate(workers):
            minDist, minBike = closestBike(iWorker)
            heapq.heappush(distHeap, (minDist, iWorker, minBike))
        # print(distHeap)
        
        wkMatches = [0]*(len(workers))
        while len(distHeap) > 0:
            _, iWorker, iBike = heapq.heappop(distHeap)
            if iBike not in bikeUsed:
                # print(iWorker, wkMatches)
                wkMatches[iWorker] = iBike
                bikeUsed.add(iBike)
            else:
                minDist, minBike = closestBike(iWorker)
                heapq.heappush(distHeap, (minDist, iWorker, minBike))
        return wkMatches

    
# 2.7sec, faster than 5%
class SolutionSlower:
    # Generate all combinations, put in min-heap, pull out matches until done
    def assignBikes(self, workers: List[List[int]], bikes: List[List[int]]) -> List[int]:
        # Computation: O(n^2); Space: O(n^2)
        distMatch = [(abs(xWk - xBk) + abs(yWk - yBk), iWk, iBk) for iWk, (xWk, yWk) in enumerate(workers) for iBk, (xBk, yBk) in enumerate(bikes)]
        # distMatch.sort() # O(n^n*log(n)) (because array length is n^2)
        heapq.heapify(distMatch)
        
        iWkFree = len(workers)
        iBkFree = len(bikes)
        unmatch = -1
        wkMatches = [unmatch]*(iWkFree)
        bkMatches = [unmatch]*(iBkFree)
        while iWkFree > 0 and iBkFree > 0:
            dist, iWk, iBk = heapq.heappop(distMatch)
            if wkMatches[iWk] == unmatch and bkMatches[iBk] == unmatch:
                wkMatches[iWk] = iBk
                bkMatches[iBk] = iWk
                iWkFree -= 1
                iBkFree -= 1
        # print("Loop done")
        return wkMatches


s = Solution()
print(s.assignBikes([[0,0],[2,1]], [[1,2],[3,3]])) # [1,0]
print(s.assignBikes([[0,0],[1,1],[2,0]], [[1,0],[2,2],[2,1]])) # [0,2,1]        



