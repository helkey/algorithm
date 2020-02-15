""" 124. Binary Tree Maximum Path Sum
For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.
https://leetcode.com/problems/binary-tree-maximum-path-sum/

Faster than 10% of Python submissions
"""

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def maxPathSum(self, root: TreeNode) -> int:
        return max(self.maxPathSums(root))

    def maxPathSums(self, node: TreeNode) -> int:
        if node == None:
            return None, 0
        updownL, upLeft = self.maxPathSums(node.left)
        updownR, upRight = self.maxPathSums(node.right)
        updown = node.val + max(upLeft, 0) + max(upRight, 0)
        updownMax = max([upDn for upDn in [updown, updownL, updownR] if upDn != None])
        upMax = node.val + max(0, upLeft, upRight)
        return updownMax, upMax

s = Solution()
n1 = TreeNode(-3)
print(s.maxPathSum(n1)) # -3

n1, n2, n3 = [TreeNode(v) for v in [1, 2, 3]]
n1.left, n1.right = n2, n3
print(s.maxPathSum(n1)) # 6

n1, n2, n3, n4, n5 = [TreeNode(v) for v in [-10, 9, 20, 15, 7]]
n1.left, n1.right = n2, n3
n3.left, n3.right = n4, n5
print(s.maxPathSum(n1)) # 42      


n1, n2, n3, n4 = [TreeNode(v) for v in [-2, 6, 0, -6]]
n1.left = n2
n2.left, n2.right = n3, n4
print(s.maxPathSum(n1)) # 6

