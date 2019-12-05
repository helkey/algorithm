/* 24. Swap Nodes in Pairs
  Given a linked list, swap every two adjacent nodes and return its head.
  You may not modify the values in the list's nodes, only nodes itself may be changed.
*/

package main

import "fmt"

func main() {
	li := &ListNode{5, nil}
        li = &ListNode{4, li}
        li = &ListNode{3, li}
        li = &ListNode{2, li}
        li = &ListNode{1, li}
        head := swapPairs(li)
	lo2, lo3, lo4 := (*head).Next, ((*head).Next).Next, (((*head).Next).Next).Next
        fmt.Println((*head).Val, (*lo2).Val, (*lo3).Val, (*lo4).Val)
}

// Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if h0 == nil || (*h0).Next == nil {
		return head
	}
	h1 := (*h0).Next
	head, h1.Next, h0.Next = h1, h1.Next, swapPairs(h1.Next)
	return head
}
