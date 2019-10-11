/* 108. Convert Sorted Array to Binary Search Tree

Array where elements are sorted in ascending order, 
convert it to a height balanced BST.

Runtime: 176 ms, faster than 50.52% of Go online submissions for Convert Sorted Array to Binary Search Tree.
Memory Usage: 163.1 MB, less than 100.00% of Go online submissions for Convert Sorted Array to Binary Search Tree.
*/

package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func main() {
	bst := sortedArrayToBST([]int{-10,-3,0,5,9})
	fmt.Println(maxDepth(bst))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	fmt.Println("sort", nums)
	iCen := len(nums)/2
	var tL *TreeNode
	if iCen > 0 {
		tL = sortedArrayToBST(nums[:iCen])
	}
	tR := sortedArrayToBST(nums[iCen+1:])
	root := TreeNode{nums[iCen], tL, tR}
	return &root
}



/** FROM 104_Maximum_Tree_Depth **/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	fmt.Println("  ", (*root).Val)
	return max(maxDepth((*root).Left), maxDepth((*root).Right)) + 1
}
func max(a, b int) int {
	if a> b {
		return a
	}
	return b
}
