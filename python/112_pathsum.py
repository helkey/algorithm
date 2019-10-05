# 112: Path Sum

# Definition for a binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        def sum_list(root: TreeNode) -> {int}:
            if not root:
                return {}
            if (not root.left) and (not root.right):
                return set([root.val])
            set_left = {s + root.val for s in sum_list(root.left)} if root.left else set()
            if not root_right:
                return set_left
            set_right = {s + root.val for s in sum_list(root.right)}
            if not root_left:
                return set_right
            return set_left.union(set_right)
        return sum in sum_list(root)

# Test case:
t1 = TreeNode(1)
t2 = TreeNode(2)
t2.right = t1

s = Solution()
print(s.hasPathSum(t2, 2)) # False
print(s.hasPathSum(t2, 3)) # True

