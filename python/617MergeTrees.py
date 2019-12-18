""" 617. Merge Two Binary Trees  https://leetcode.com/problems/merge-two-binary-trees
Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Faster than 92% of Python submissions
"""

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
        def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
            if t1 != None and t2 != None:
                t1.val += t2.val                
                t1.left = self.mergeTrees(t1.left, t2.left)
                t1.right = self.mergeTrees(t1.right, t2.right)
            if t1 == None:
                t1 = t2
            return t1

        
    
