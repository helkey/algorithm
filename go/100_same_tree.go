/* 100. Same Tree
Given two binary trees, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical and the nodes have the same value. */
// faster than 100.00% of Go online submissions for Same Tree.

package main


// Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pNil := p == nil
	qNil := q == nil
	if pNil && qNil {
		return true
	}
	if pNil || qNil {
		return false
	}
	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
