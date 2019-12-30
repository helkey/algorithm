""" 844. Backspace String Compare
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.
  https://leetcode.com/problems/backspace-string-compare/

Faster than 99% of Python submissions
"""

class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:
        def charYield(S: str) -> str:
            sRev = S[::-1] # reverse string
            skip = 0
            for c in sRev:
                if c == "#":
                    skip += 1
                elif skip > 0:
                    skip -= 1
                else:
                    yield c
        yieldS = charYield(S)
        yieldT = charYield(T)
        for sChar in yieldS:
            try:
                tChar = next(yieldT)
                # print(sChar, tChar)
                if sChar != tChar:
                    # print("Unequal char")
                    return False
            except:
                # print("T is shorter string")
                return False
        
        # Check if yieldT empty
        try:
            tChar = next(yieldT)
            # print("T is longer string")
            return False # different length strings
        except:
            return True # ALL characters matched
        
s = Solution()
print(s.backspaceCompare("ab#c", "ad#c")) # True
print(s.backspaceCompare("ab##", "c#d#")) # True
print(s.backspaceCompare("a##c", "#a#c")) # True
print(s.backspaceCompare("a#c", "#b")) # False

