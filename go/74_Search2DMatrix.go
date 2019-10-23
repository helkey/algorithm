/* 74. Search a 2D Matrix

O(log(n)) binary search
(faster than 98% of Go submissions) */

package main
import "fmt"

func main() {
	matrix := [][]int{[]int{1,   3,  5,  7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50}}
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
	nrow := len(matrix)
	if nrow == 0  || (nrow == 1 && len(matrix[0]) == 0) {
		return false
	}

	ncol := len(matrix[0])
	if (target < matrix[0][0]) || (target > matrix[nrow-1][ncol-1]) {
		return false
	}
	if (target == matrix[0][0]) || (target == matrix[nrow-1][ncol-1]) {
		return true
	}
	nTot := nrow * ncol
	iStart := 0
	//iEnd := nTot - 1
	for iEnd := nTot - 1; iEnd - iStart > 1; {
		iMid := (iStart + iEnd) / 2
		vMid := getVal(matrix, iMid, nrow, ncol)
		if vMid == target {
			return true
		}
		if vMid > target {
			iEnd = iMid
		} else {
			iStart = iMid
		}
		// fmt.Println(iMid, vMid)
	}
	return false
}

func getVal(matrix [][]int, indx, nrow, ncol int) int {
	irow := indx / ncol
	icol := indx % ncol
	return matrix[irow][icol]
}
	
