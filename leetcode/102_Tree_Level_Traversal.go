/** 102. Binary Tree Level Order Traversal

**/

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main () {
	b1 := TreeNode{9, nil, nil}
	b2 := TreeNode{20, nil, nil}
	root := TreeNode{3, &b1, &b2}
	fmt.Println(levelOrder(&root))
}

func levelOrder(root *TreeNode) [][]int {
	new_order := [][]int{}
	return *nextLevel(root, &new_order, 0)
	// fmt.Println("  Len", len(order))
}

func nextLevel(root *TreeNode, order *[][]int, level int) *[][]int {
	if root != nil {
		if len(*order) < level + 1 {
			new_order := append(*order, []int{})
			order = &new_order
		}
		row := &(*order)[level]
		if len(*row) == 0 {
			*row = []int{root.Val}
		} else {
			*row = append(*row, root.Val)
		}
		// fmt.Println(level, (*order)[0], *row)
		order = nextLevel(root.Left, order, level + 1)
		order = nextLevel(root.Right, order, level + 1)
		// fmt.Println("  Len", len(*order))
	}
	return order
}
