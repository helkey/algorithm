/** 118. Pascal's Triangle

Generate the first numRows of Pascal's triangle.
e.g. n=5
[    [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1] ]

Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Pascal's Triangle.
**/

package main

import "fmt"

func main() {
	pt := generate(5)
	for _, r := range pt {
		fmt.Println(r)
	}
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	pt := [][]int{[]int{1}}
	if numRows == 1 {
		return pt
	}
	pt = append(pt, []int{1, 1})
	rowLast :=  pt[1]
	for n:=3; n<=numRows; n++ {
		row :=make([]int, n)
		row[0] = 1
		row[n-1] = 1
		for r:=1; r<n-1; r++ {
			row[r] = rowLast[r-1] + rowLast[r]
		}
		pt = append(pt,  row)
		rowLast = row
	}
	return pt
}

