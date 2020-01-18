""" 235. Lowest Common Ancestor of a Binary Search Tree
Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.
Definition of LCA on Wikipedia: The lowest common ancestor is defined between two nodes p and q as the lowest node in T
  that has both p and q as descendants (where we allow a node to be a descendant of itself).
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

THIS is BINARY search tree; compare root.val, p, q to know which side to search
  https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/discuss/467827/Java-Solution-using-Recursion-4ms-faster-than-100

Faster than 80% of Python submissions
"""

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        def pqInNode(root, p, q):
            if root == None:
                return None
            if root.val < p:
                return pqInNode(root.right, p, q)
            if root.val > q:
                return pqInNode(root.left, p, q)
            return root

        if p.val > q.val:
            p, q = q, p # make sure target nodes in ascending order
        node = pqInNode(root, p.val, q.val)
        return node
    

s = Solution()            
root, n2, n8, n0, n4, n7, n9, n3, n5 = [TreeNode(v) for v in [6,2,8,0,4,7,9,3,5]]
root.left, root.right = n2, n8
n2.left, n2.right, n8.left, n8.right = n0, n4, n7, n9
n4.left, n4.right = n3, n5
print((s.lowestCommonAncestor(root, n2, n8)).val) # 6
print((s.lowestCommonAncestor(root, n2, n4)).val) # 2
