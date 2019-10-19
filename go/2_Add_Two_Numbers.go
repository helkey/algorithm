/* 2. Add Two Numbers
   Given two non-empty linked lists representing two non-negative integers, with digits stored in reverse order,
     add the two numbers and return it as a linked list.
(Faster than 90% of Go submissions) */

package main

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, sum_last *ListNode
	for carry:=0; (l1 != nil) || (l2 != nil) || (carry != 0); {
		x1, x2 := 0, 0
		if (l1 != nil) {
			x1 = l1.Val
			l1 = l1.Next
		}
		if (l2 != nil) {
			x2 = l2.Val
			l2 = l2.Next
		}
		sum := &ListNode{x1 + x2 + carry, nil}
		if (sum.Val > 9) {
			sum.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		if head == nil {
			head = sum
		} else {
			sum_last.Next = sum
		}
		sum_last = sum
	}
	return head
}
