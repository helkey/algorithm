/** 104. Maximum Depth of Binary Tree

Maximum depth is the number of nodes along the longest path 
from the root node down to the farthest leaf node.

e.g depth = 3:
    3
   / \
  9  20
    /  \
   15   7

Runtime: 8 ms, faster than 71.77% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.4 MB, less than 100.00% of Go online submissions for Maximum Depth of Binary Tree.
**/

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func main() {
	r2 := TreeNode{2, nil, nil}
	r1 := TreeNode{1, &r2, nil}
	r0 := TreeNode{0, nil, &r1}
	fmt.Println(maxDepth(&r0))
}
	
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth((*root).Left), maxDepth((*root).Right)) + 1
}

func max(a, b int) int {
	if a> b {
		return a
	}
	return b
}
