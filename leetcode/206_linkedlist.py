#206. Reverse Linked List
# Faster than 83.30% of Python3 online submissions

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head == None:
            return head
        last = None
        while head.next != None:
            next = head.next
            head.next = last
            last = head
            head = next
            # Why doesn't this work?
            # head, head.next, last = head.next, last, head
        head.next = last
        return head

    
a = ListNode(1); b = ListNode(2); c = ListNode(3)
a.next = b; b.next = c
print(a.val, a.next.val, a.next.next.val)

s = Solution()
x = s.reverseList(a)
print(x.val)
print(a.val, a.next, b.val, b.next.val, c.val, c.next.val)

