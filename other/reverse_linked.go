// reverse_linked.go: Reverse linked list
package main

import "fmt"


type element struct {
  val       int
  next      *element
}


func main() {
	head := genLL([]int{1, 6, 2, 3})
	printLL(head)
	head = revLL(head)
	printLL(head)
}

// reverse linked list
func revLL(head *element) (last *element) {
	// last is initialized to nil
	for h := head; h != nil;  {
		// 3-way variable swap
		h, (*h).next, last = (*h).next, last, h
	}
	return
}
		
// Initialize linked list from array
func genLL(arr []int) (head *element) {
	for i := len(arr)-1; i >= 0; i--  {
		e := new(element)
		e.val = arr[i]
		e.next = head
		head = e
	}
	return
}

// Print linked list
func printLL(head *element) {
	for h := head; h != nil; h=(*h).next {
		fmt.Print((*h).val, ", ")
	}
	fmt.Println("]")
}
