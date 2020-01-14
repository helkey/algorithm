""" 98. Validate Binary Search Tree
Given a binary tree, determine if it is a valid binary search tree (BST).
Assume a BST is defined as follows:
  *The left subtree of a node contains only nodes with keys less than the node's key.
  *The right subtree of a node contains only nodes with keys greater than the node's key.
  *Both the left and right subtrees must also be binary search trees.
  https://leetcode.com/problems/validate-binary-search-tree/

Faster than 62% of Python submissions
"""

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        def valMinMax(root) -> (int, int):
            if root == None: # or (root.left == None and root.right == None):
                return True, None, None # leaf nodes are balanced
            if root.left != None:
                balL, minL, maxL =  valMinMax(root.left)
                gtLeft = root.val > maxL
            else:
                balL, gtLeft, minL = True, True, root.val
            if root.right != None:
                balR, minR, maxR =  valMinMax(root.right)
                ltRight = root.val < minR
            else:
                balR, ltRight, maxR = True, True, root.val
            isBal = balL and balR and gtLeft and ltRight
            return isBal, minL, maxR

        isValid, _, _ = valMinMax(root)
        return isValid
        
s = Solution()
vals = [2,1,3] # true
root, n1, n3 = [TreeNode(v) for v in vals]
root.left, root.right = n1, n3
print(s.isValidBST(root))

vals = [5,1,4,None,None,3,6] # false
root, n1, n4, n0, n0, n3, n6 = [TreeNode(v) for v in vals]
root.left, root.right = n1, n4
n4.left, n4.right = n3, n6
print(s.isValidBST(root))

vals = [1,1]
root, n1 = [TreeNode(v) for v in vals]
root.right = n1
print(s.isValidBST(root))
