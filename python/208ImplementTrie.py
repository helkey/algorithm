""" 208. Implement Trie (Prefix Tree)
Implement a trie with insert, search, and startsWith methods.
  https://leetcode.com/problems/implement-trie-prefix-tree/

Faster than 85% of Python submisisons
"""

# from collections import defaultdict

class Trie:
    def __init__(self):
        """ Initialize your data structure here.
        """
        self.end_char = "#"
        self.trie = {}

    def insert(self, word: str) -> None:
        """ Inserts a word into the trie.
        """
        t = self.trie
        for c in word:
            if c not in t:
                t[c] = {}
            t = t[c]
        t[self.end_char] = None
            

    def search(self, word: str) -> bool:
        """ Returns if the word is in the trie.
        """
        t = self.trie
        for c in word:
            if c not in t:
                return False
            t = t[c]
        if self.end_char not in t:
            return False
        return True
    
    def startsWith(self, prefix: str) -> bool:
        """ Returns if there is any word in the trie that starts with the given prefix.
        """
        t = self.trie
        for c in prefix:
            if c not in t:
                return False
            t = t[c]
        return True
        

trie = Trie()
trie.insert("apple")
print(trie.search("apple")) # -> True
print(trie.search("app")) # -> False
print(trie.startsWith("app")) # -> True
print(trie.startsWith("apple")) # -> True
trie.insert("app")
print(trie.search("app")) # -> True

