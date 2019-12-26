""" 1170. Compare Strings by Frequency of the Smallest Character
Define function f(s) over a non-empty string s, calculates the frequency of the smallest character in s. For example, if s = "dcce" then f(s) = 2 because the smallest character is "c" and its frequency is 2.
Now, given string arrays queries and words, return an integer array answer, where each answer[i] is the number of words such that f(queries[i]) < f(W), where W is a word in words.
  https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/

Faster than 90% of Python submissions
"""

from bisect import bisect
from typing import List

class Solution:
    def numSmallerByFrequency(self, queries: List[str], words: List[str]) -> List[int]:
        wordFreqs = sorted([self.freq(word) for word in words])
        queryFreqs = [self.freq(q) for q in queries]
        # return [self.count(queryFreq, wordFreqs) for queryFreq in queryFreqs]
        return [len(words) - bisect(wordFreqs, self.freq(query)) for query in queries]

    """ FASTER to use s.count(min(s)) """
    def freq(self, s) -> int:
        """ Frequency of the smallest char in s """
        cnt = 1
        minO = ord(s[0])
        for c in s[1:]:
            o = ord(c)
            if o == minO:
                cnt += 1
            elif o < minO:
                minO = o
                cnt = 1
        return cnt

    """ USE faster 'bisect' library """
    def count(self, queryFreq: int, wordFreq: List[int] ) -> int:
        """ Count query < word for how many words in (sorted) wordFreq """
        if queryFreq < wordFreq[0]:
            return len(wordFreq)
        if queryFreq >= wordFreq[-1]:
            return 0
        ptrSt, ptrEnd = 0, len(wordFreq) - 1
        while ptrEnd - ptrSt > 1:
            ptrM = (ptrSt + ptrEnd) // 2 # center of search range
            # print(queryFreq, ptrSt, ptrM, ptrEnd, wordFreq[ptrM])
            if queryFreq < wordFreq[ptrM]:
                ptrEnd = ptrM
            else:
                ptrSt = ptrM
        return len(wordFreq) - ptrSt - 1
                
s = Solution()
print(s.numSmallerByFrequency(["cbd"], words = ["zaaaz"])) # assert [1]
print(s.numSmallerByFrequency(["bbb","a", "cc", "dddd"], words = ["a","aa","aaa","aaaa"])) # assert [1, 2]
