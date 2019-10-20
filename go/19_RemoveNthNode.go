/* 19. Remove Nth Node From End of List
   Given a linked list, remove the n-th node from the end of list and return its head.

leetcode.com/problems/remove-nth-node-from-end-of-list */

package main

import "fmt"

// Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
}

func main() {
	ln := &ListNode{5, nil}
	ln = &ListNode{4, ln}
	ln = &ListNode{3, ln}
	ln = &ListNode{2, ln}
	ln = &ListNode{1, ln}
	lO := removeNthFromEnd(ln, 5)
	fmt.Println((*lO).Val, (*(*lO).Next).Val)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ptr1 = head
	var ptr2 = head
	var ptr2Last *ListNode
	for i:=0; (*ptr1).Next!=nil; i++ {
		ptr1 = (*ptr1).Next
		// fmt.Println((*ptr1).Val)
		// Start ptr1 n steps ahead of ptr2
		if i >= n - 1 {
			ptr2Last = ptr2
			ptr2 = (*ptr2).Next
		}
	}
	// ptr2 pointing to N'th node
	fmt.Println((*ptr2).Val)
	if ptr2Last == nil {
		return head.Next
	}
	(*ptr2Last).Next = (*ptr2).Next
	return head
}
