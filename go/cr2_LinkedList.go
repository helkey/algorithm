

package main

import (
       "fmt"
)

type Linked_list struct {
     len int
     start *listElem
     end *listElem
}

type listElem struct {
     i int
     next *listElem
}

func main() {
     test_list()
}

func (ll *Linked_list) Push(i int) {
           elem_new := listElem{}
           elem_new.i = i
           ll.end = &elem_new
           ll.len++
}

func (ll *Linked_list) Pop() int {
      return ll.end.i
}

func test_list() {
    l := Linked_list{}
    l.Push(314)
    i := l.Pop()
    fmt.Println(i)
}