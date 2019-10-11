/** 509. Fibonacci

Fibonacci sequence: each number is the sum of the two preceding ones, 
starting from 0 and 1

Runtime: 0 ms, faster than 100.00% of Go online submissions for Fibonacci Number.
Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Fibonacci Number.
**/

package main

import "fmt"

func main() {
	for n:=0; n<5; n++ {
		fmt.Println(n, fib(n)) // 0, 1, 1, 2, 3...
	}
}

func fib(n int) int {
	seq := []int{0, 1}
	for i:=2; i<=n; i++ {
		lenS := len(seq)
		seq = append(seq, seq[lenS-2] + seq[lenS-1])
	}
	return seq[n]
}
