# 344. Reverse String
# Faster than 93% of Python3 submissions

from typing import List

class Solution:
    def reverseString(self, s: List[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        ln = len(s)
        len2 = int(ln/2)
        for i in range(len2):
            o = ln - i - 1
            s[i], s[o] = s[o], s[i]

s= Solution()
st = ["h","e","l","l","o"]
s.reverseString(st)
print(st)

st = ["a", "b", "c", "d"]
s.reverseString(st)
print(st)



