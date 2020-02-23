""" Find the Celebrity
Suppose you are at a party with n people (labeled from 0 to n - 1) and among them, there may exist one celebrity. 
The definition of a celebrity is that all the other n - 1 people know him/her but he/she does not know any of them.
https://leetcode.com/problems/find-the-celebrity/

Faster than 65% of Python submissions
"""

class Solution:
    def findCelebrity(self, n: int) -> int:
        # Identify single candidate as possible celebrity
        iStart, iEnd = 0, n - 1
        while iEnd - iStart > 0:
            if knows(iStart, iEnd):
                iStart += 1
            else:
                iEnd -= 1

        # Everybody knows celebrity; celebrity doesn't know anybody
        # celebDoesntKnow = not any([knows(iStart, i) for i in range(n) if i != iStart])
        for i in range(n):
            if i != iStart and not knows(i, iStart):
                return -1 # somebody doesn't know celebrity candidate
            if i != iStart and knows(iStart, i):
                return -1 # celebrity candidate knows someone else
        return iStart
    
s = Solution()
graph1 = [[1,1,0],
         [0,1,0],
         [1,1,1]] # 1

graph = [[1,0,1],
         [1,1,0],
         [0,1,1]] # -1

graph = [[1,1],[1,1]] # -1

def knows(i1, i2) -> bool:
    # iCeleb = 8
    # return (i2 == iCeleb)
    return graph[i1][i2]

print(s.findCelebrity(2))

