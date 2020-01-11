""" 1249. Minimum Remove to Make Valid Parentheses
Given a string s of '(' , ')' and lowercase English characters. 
Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) 
  so that the resulting parentheses string is valid and return any valid string.
A parentheses string is valid if and only if:
  *It is the empty string, contains only lowercase characters, or
  *It can be written as AB (A concatenated with B), where A and B are valid strings, or
  *It can be written as (A), where A is a valid string.
https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

"""

class Solution:
    """ Process forward, remove all unmatched closing brackets (output to list for O(N))
        Process backwards, remove all unmatched opening brackets """
    def minRemoveToMakeValid(self, s: str) -> str:
        def trimClosing(s, openC, closeC: str) -> str:
            cnt = 0
            out = []
            for c in s:
                if cnt > 0 or c != closeC:
                    out.append(c)
                    if c == closeC:
                        cnt -= 1
                if c == openC:
                    cnt += 1
            out.reverse()
            return "".join(out)
        # trim extra closing chars ")"
        sClose = trimClosing(s, "(", ")")
        # reverse string, trim extra opening chars "("
        sOpen = trimClosing(sClose, ")", "(")
        return sOpen
                            
                
s = Solution()
print(s.minRemoveToMakeValid("lee(t(c)o)de)")) # "lee(t(c)o)de"
print(s.minRemoveToMakeValid("a)b(c)d")) # "ab(c)d"
print(s.minRemoveToMakeValid("))((")+"#") # ""
print(s.minRemoveToMakeValid("(a(b(c)d)")) # "a(b(c)d)"

