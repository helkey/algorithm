""" 253. Meeting Rooms II
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), 
  find the minimum number of conference rooms 

Faster than 85% of Python submissions
"""

from typing import List

class Solution:
    # Keep track of all starting/stopping times; don't need linkage between them
    # O(n*logn) for sorting start and stop time arras
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        if len(intervals) <= 1:
            return len(intervals)
        startT, stopT = [list(tArr) for tArr in list(zip(*intervals))]
        startT.sort()
        stopT.sort()
        rooms, maxrooms, iStop = 0, 0, 0
        for tStart in startT:
            while stopT[iStop] <= tStart:
                # Deallocate concluded meetings
                rooms -= 1
                iStop += 1
            rooms += 1 # allocate new meeting
            maxrooms = max(rooms, maxrooms)
        return maxrooms

s = Solution()
print(s.minMeetingRooms([[0, 30],[5, 10],[15, 20]])) # 2
print(s.minMeetingRooms([[7,10],[2,4]])) #1
