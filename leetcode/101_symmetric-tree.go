/* 101. Symmetric Tree
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center). */
// Faster than 100% of Go online submissions for Symmetric Tree.

package main

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}
