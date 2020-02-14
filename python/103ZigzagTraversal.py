""" 103. Binary Tree Zigzag Level Order Traversal
Given a binary tree, return the zigzag level order traversal of its nodes' values. 
  (ie, from left to right, then right to left for the next level and alternate between).
  https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

"""

from collections import dequeue
from typing import List

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        traverse = []
        if root == None:
            return traverse
        row = [root]
        cnt = 0
        while len(row) > 0:
            if cnt % 2 == 0:
                rowVal = [node.val for node in row]
            else:
                rowVal = [node.val for node in row[::-1]]
            traverse.append(rowVal)
            row_next = dequeue()
            for node in row:
                if node.left:
                    row_next.append(node.left)
                if node.right:
                    row_next.append(node.right)
            row = row_next
            cnt += 1
        return traverse
    
    # Faster than 6%!
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        traverse = []
        if root == None:
            return traverse
        row = [root]
        cnt = 0
        while len(row) > 0:
            if cnt % 2 == 0:
                rowVal = [node.val for node in row]
            else:
                rowVal = [node.val for node in row[::-1]]
            traverse.append(rowVal)
            row = [child for node in row for child in [node.left, node.right] if child != None]
            cnt += 1
        return traverse

s = Solution()
root, n9, n20, n15, n7 = [TreeNode(n) for n in [3, 9, 20, 15, 7]]
root.left, root.right = n9, n20
n20.left, n20.right = n15, n7
print(s.zigzagLevelOrder(root))


            

            
