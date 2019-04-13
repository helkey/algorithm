/** 23. Merge k sorted lists
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
**/

package main

import "fmt"

type ListNode struct {
	 Val int
	 Next *ListNode
}

func main() {
	node1 := createList([]int{1, 4, 5})
	node2 := createList([]int{1, 3, 4})
	// node3 := createList([]int{2, 6}) // works
	// node3 := &ListNode{2, nil} // works
	node3 := &ListNode{} // works
	node := mergeKLists([]*ListNode{node1, node2, node3})
	for node != nil {
		fmt.Print(node.Val, " ")
		node = (*node).Next
	}
	fmt.Println()
}		

func mergeKLists(lists []*ListNode) *ListNode {
	node_head := ListNode{}
	node_last := &node_head
	var i_min int
	var min int
	for {
		// Find next min value
		i_min = -1
		for i, node := range lists {
			if (node != nil) && ((i_min == -1) || (node.Val < min)) {
				i_min = i
				min = node.Val
			}
		}
		if i_min == -1 {
			return node_head.Next
		}
		node_new := ListNode{min, nil}
		node_last.Next = &node_new
		node_last = &node_new
		// Remove minimum ListNode from list
		node_min := lists[i_min]
		if node_min.Next != nil {
			// Select next element in list
			lists[i_min] = node_min.Next
		} else {
			// Remove empty list (PROBABLY DONT NEED TO DO!)
			lists = append(lists[:i_min], lists[i_min+1:]...)
		}
	}
}

func createList(list []int) *ListNode {
	node := ListNode{list[len(list)-1], nil}
	for i := len(list) - 2; i >= 0; i-- {
		nodeLast := node
		node = ListNode{list[i], &nodeLast}
	}
	return &node
}
