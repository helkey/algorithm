// 110.


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isBalanced(self, root: TreeNode) -> bool:
        def treeDepth(node: TreeNode) -> int:
            if (node == None):
                return 0
            depL = treeDepth(node.left)
            depR = treeDepth(node.right)
            isBal = ((depL <= depR + 1) and (depR <= depL + 1))
            if not isBal:
                raise Exception
            return max(depL, depR) + 1
        try:
            _ = treeDepth(root)
            return True
        except:
            return False


""" SCALA: Similar implementatin blows up stack
object Solution {
  def isBalanced(root: TreeNode): Boolean = {
    val (_, balanced) = treeDepth(root)
    return balanced
    }
  
  def treeDepth(node: TreeNode): (Int, Boolean) = {
    if (node == null) {
      return (0, true)
    }
    val (depL, balL) = treeDepth(node.left)
    val (depR, balR) = treeDepth(node.right)
    val bal = ((depL <= depR + 1) && (depR <= depL + 1)) // this node is balanced
    return (List(depL, depR).max + 1, (bal && balL && balR))
   }
} """

""" This approach checks for minimum height tree;
    which does not satisfy problem test case:
      Input: [1,2,2,3,3,3,3,4,4,4,4,4,4,null,null,5,5]
      Output: False; Expected: True
object Solution {
  def isBalanced(root: TreeNode): Boolean = {
    val (minD, maxD) = minMaxDepth(root)
    return (maxD <= minD + 1)
    }
  
  def minMaxDepth(node: TreeNode): (Int, Int) = {
    if (node == null) {
      return (0, 0)
    }
    val (minL, maxL) = minMaxDepth(node.left)
    val (minR, maxR) = minMaxDepth(node.right)
    return (List(minL, minR).min + 1, List(maxL, maxR).max + 1)
   }
} """
