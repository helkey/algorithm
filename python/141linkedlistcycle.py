""" 141. Linked List Cycle  

Given a linked list, determine if it has a cycle in it.
  https://leetcode.com/problems/linked-list-cycle/

Faster than 98% of Python submissions
"""

# from typing import List, ...

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

# Use two pointers traveling at rates x, 2x
class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head == None:
            return False
        ptrX = head
        ptr2X = head
        while True:
            if not ptr2X.next or not (ptr2X.next).next:
                return False
            ptrX = ptrX.next
            ptr2X = (ptr2X.next).next
            if ptrX == ptr2X:
                return True

s = Solution()

print(s.hasCycle(None)) # True

h1, h2, h3, h4 = ListNode(3), ListNode(2), ListNode(0), ListNode(-4)
h1.next, h2.next, h3.next, h4.next = h2, h3, h4, h2
print(s.hasCycle(h1)) # True

h1, h2 = ListNode(1), ListNode(2)
h1.next, h2.next = h2, h1
print(s.hasCycle(h1)) # True

h2.next = None
print(s.hasCycle(h1)) # False
