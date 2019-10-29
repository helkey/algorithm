/* 841. Keys and Rooms

(Faster than 94% of Go submissions) */

package main

import (
	"fmt"
)

func main() {
	rooms := [][]int{{1}, {2}, {3}, {}} // true
	// rooms := [][]int{{1,3}, {3,0,1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms))
}


// Uses list/stack to stack rooms, keys
func canVisitAllRooms(rooms [][]int) bool {
	nRooms := len(rooms)
	openRooms := make([]bool, nRooms)
	keys := Stack()
	keys.Push(0) // start with key to room 0
	
	for ; keys.Len()>0; {
		roomKey := keys.Pop()
		if !openRooms[roomKey] {
			for _, key := range rooms[roomKey] {
				if !openRooms[key] {
					keys.Push(key)
				}
			}
		}
		openRooms[roomKey] = true
		// fmt.Println("Open", openRooms, keys.Len())
	}
	for _, open := range openRooms {
		if !open {
			return false
		}
	}
	return true		
}

type stack struct {
      arr []int
}

func Stack() *stack {
	return &stack {make([]int, 0)}
} 

func (s *stack) Len() int {
	return len(s.arr)
}

func (s *stack) Push(v int) {
	s.arr = append(s.arr, v)
}

func (s *stack) Pop() int {
	ln := len(s.arr)
	res := s.arr[ln - 1]
	s.arr = s.arr[:ln - 1]
	return res
}



/* Faster than 28% of Go submissions
// Uses maps to track rooms/keys
func canVisitAllRooms(rooms [][]int) bool {
	// fmt.Println(rooms)
	keys := map[int]bool{}
	// Get the keys out of room 0
	for _, key := range rooms[0] {
		keys[key] = true
	}
	// Mark all other rooms locked
	locked := map[int]bool{}
	for i:=1; i<len(rooms); i++ {
		locked[i] = true
	}
	for ; ; {
		if len(locked) == 0 {
			return true // all rooms unlocked
		}
		if len(keys) == 0 {
			return false // no keys left to try
		}
		for key, _ := range keys {
			if locked[key] {
				delete(locked, key) // unlock room
				// get keys from this room
				for _, newKey := range rooms[key] {
					keys[newKey] = true
				}
			}
			delete(keys, key)
		}
	}
} */
