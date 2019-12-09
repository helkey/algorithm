""" 819. Most Common Word

Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  
  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.
  Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive. 
  Answer is in lowercase.

Fast than 88% of Python solutions
"""

import collections
import string
from typing import List

# Problem is  trivial - hard part is extracting punctionation 
class Solution:
    def mostCommonWord(self, paragraph: str, banned: List[str]) -> str:
        ban = set(banned)
        wordCount = collections.defaultdict(int)
        paragraph = paragraph.replace(",", " ") # some words have no space after comma
        paragraph = paragraph.replace("  ", " ")
        paragraph = paragraph.translate(str.maketrans('', '', string.punctuation))
        words = paragraph.lower().split(" ")
        for word in words:
            if word not in ban:
                wordCount[word] += 1
        maxCnt = 0
        match = ""
        for word in wordCount:
            count = wordCount[word]
            if count > maxCnt:
                maxCnt = count
                match = word
        return match
    
        

s = Solution()
print(s.mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", ["hit"])) # ball
print(s.mostCommonWord("a, a, a, a, b,b,b,c, c", ["a"])) # b
