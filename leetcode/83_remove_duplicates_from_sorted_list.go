// 83. Remove Duplicates from Sorted List
//   Given a sorted linked list, delete all duplicates such that each element appear only once.
//  Faster than 90% of Go online submissions for Remove Duplicates from Sorted List.

package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
 }
 
func main () {
	head := genLL([]int{1,1,2})
	head = deleteDuplicates(head)
	printLL(head)
	head = genLL([]int{1,1,2,3,3})
	head = deleteDuplicates(head)
	printLL(head)
	head = genLL([]int{1,1,1})
	head = deleteDuplicates(head)
	printLL(head)
	head = genLL([]int{})
	head = deleteDuplicates(head)
	printLL(head)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for h:=head; h.Next != nil;  {
		if (h.Val == (h.Next).Val) {
			h.Next = (h.Next).Next
		} else {
			h = h.Next
		}
	}
	return head
}


// Initialize linked list from array
func genLL(arr []int) (head *ListNode) {
	for i := len(arr)-1; i >= 0; i--  {
		e := new(ListNode)
		e.Val = arr[i]
		e.Next = head
		head = e
	}
	return
}

// Print linked list
func printLL(head *ListNode) {
	for h := head; h != nil; h=(*h).Next {
		fmt.Print((*h).Val, ", ")
	}
	fmt.Println("]")
}

