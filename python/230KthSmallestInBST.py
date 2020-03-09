"""230. Kth Smallest Element in a BST
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
  https://leetcode.com/problems/kth-smallest-element-in-a-bst/

Faster than 98% of Python submissions
"""

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        i = 0
        kthVal = None
        def DFS(node: TreeNode, k):
            nonlocal i, kthVal
            if (node == None) or (i > k):
                return
            DFS(node.left, k)
            i += 1
            # print(i, k, node.val)
            if i == k:
                kthVal = node.val
            DFS(node.right, k)
        DFS(root, k)
        return kthVal
            
s = Solution()

n3, n1, n4, n2 = [TreeNode(i) for i in [3, 1, 4, 2]]
n3.left, n3.right = n1, n4
n1.right = n2
print(s.kthSmallest(n3, 1)) # assert 1

n5, n3, n6, n2, n4, n1 = [TreeNode(i) for i in [5, 3, 6, 2, 4, 1]]
n5.left, n5.right = n3, n6
n3.left, n3.right = n2, n4
n2.left = n1
print(s.kthSmallest(n5, 3)) # assert 3
