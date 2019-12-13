""" 692. Top K Frequent Words
Given a non-empty list of words, return the k most frequent elements.
  Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the 
  lower alphabetical order comes first. https://leetcode.com/problems/top-k-frequent-words/

Faster than 95% of Python submissions

Follow-up: solve it in O(n log k) time and O(n) extra space 
  (don't need full sort to find highest K values!)
"""

from collections import defaultdict
from typing import List

class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        freq = defaultdict(int)
        for word in words:
            freq[word] += 1
        # Two-level sort: -l[0] for reverse sort, l[1] for secondary alphabetical sort
        wordfreq = sorted([[freq[word], word] for word in freq],
                          key = lambda l:(-l[0], l[1]))
        return [word for _, word in wordfreq[:k]]
        

s = Solution()
print(s.topKFrequent(["i", "love", "leetcode", "i", "love", "coding"], 2))
