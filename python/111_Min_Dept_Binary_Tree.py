# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def minDepth(self, root: TreeNode) -> int:
        if not root: return 0
        def mindp(node):
            if not node: return 0
            if node.left and node.right:
                return min(mindp(node.left)+1, mindp(node.right)+1)
            if node.left: return mindp(node.left)+1
            if node.right: return mindp(node.right)+1
            return 1
        return mindp(root)
    
