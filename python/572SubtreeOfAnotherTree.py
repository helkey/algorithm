""" 572. Subtree of Another Tree
Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.
  https://leetcode.com/problems/subtree-of-another-tree/

Faster than 33% of Python submissions
"""



# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isSubtree(self, s: TreeNode, t: TreeNode) -> bool:
        if s == None and t == None:
            return True
        if s == None or t == None:
            return False
        if s.val == t.val and self.isEqual(s.left, t.left) and self.isEqual(s.right, t.right):
            return True
        return self.isSubtree(s.left, t) or self.isSubtree(s.right, t)
    
    def isEqual(self, s: TreeNode, t: TreeNode) -> bool:
        if s == None and t == None:
            return True
        if s == None or t == None:
            return False
        return s.val == t.val and self.isEqual(s.left, t.left) and self.isEqual(s.right, t.right)
    
