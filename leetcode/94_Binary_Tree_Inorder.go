/** 94. Binary Tree Inorder Traversal

Given a binary tree, return the inorder traversal of its nodes' values.

**/

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	left := TreeNode{1, nil, nil}
	right := TreeNode{3, nil, nil}
	root := TreeNode{2, &left, &right}
	fmt.Println(inorderTraversal(&root))
}
	
func inorderTraversal(root *TreeNode) []int {
	trav := []int{}
	if root == nil {
		return trav
	}
	trav = inorderTraversal((*root).Left)
	trav = append(trav, (*root).Val)
	trav = append(trav, inorderTraversal((*root).Right)...)
	return trav
}
