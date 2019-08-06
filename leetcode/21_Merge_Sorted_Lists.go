/** 21. Merge Two Sorted Lists

Merge two sorted linked lists and return it as a new list. 
The new list should be made by splicing together the nodes of the first two lists.

Runtime: 4 ms, faster than 98.43% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Merge Two Sorted Lists.

**/

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func main() {
	mergeTwoLists(makeL([]int{1, 2, 4}), makeL([]int{1, 3, 4})) // 1, 1, 2, 3, 4, 4
	// mergeTwoLists(makeL([]int{1, 2, 4}), makeL([]int{})) // 1, 1, 2, 3, 4, 4
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	last := &head
	var minL1 bool
	for (l1 != nil) || (l2 != nil) {
		if l1 == nil {
			minL1 = false
		} else if l2 == nil {
			minL1 = true
		} else {
			minL1 = (l1.Val < l2.Val)
		}
		if minL1 {
			(*last).Next = l1
			last = l1
			l1 = (*l1).Next
		} else {
			(*last).Next = l2
			last = l2
			l2 = (*l2).Next
		}
	}
	printL(head.Next)
	return head.Next
}


func printL(l *ListNode) {
	for l != nil {
		fmt.Print((*l).Val, " ")
		l = (*l).Next
	}
	fmt.Println()
}

func makeL(arr []int) *ListNode {
	head := ListNode{}
	nLast := &head
	for _, a := range arr {
		n := ListNode{a, nil}
		(*nLast).Next = &n
		nLast = &n
	}
	return head.Next
}
			
