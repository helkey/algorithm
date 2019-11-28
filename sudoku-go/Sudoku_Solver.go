/** Sudoku Solver
Digits 1-9 must occur exactly once in each row.
  Each of the digits 1-9 must occur exactly once in each column.
  Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
**/

package main

import "fmt"

const nXY = 9 // length of row/cols
const nSect = 3 // len/number of subsections
const nBoard = nXY*nXY

func main() {
	board := testBoard()
	printBoard(board)
	solveSudoku(board)
	printBoard(board)
}

// Set up for recursive solver
func solveSudoku(board [][]byte) bool  {
	openSq := findOpenSquares(&board)
	openVals := findOpenVals(board)
	solved := iterSudoku(&board, &openSq, &openVals, 0)
	if solved { fmt.Println("SOLVED") } else {
		fmt.Println("NOT solved") }
	return solved
}

// Select one value, iterate/backtrack
func iterSudoku(board *[][]byte, openSq *[]int, openVals *openVal, iter int) bool {
	if iter == len(*openSq) {
		// Found valid entry for every square
		return true
	}
	iSqr := (*openSq)[iter] // square number to fill
	xys := x_y_s_vals[iSqr]
	vals := findValidVals(iSqr, openVals)
	for _, val := range vals {
	        // Selected val no longer available in the row/col/sect
		delete(openVals.xval[xys.x], val)
		delete(openVals.yval[xys.y], val)
		delete(openVals.sval[xys.s], val)
		(*board)[xys.y][xys.x] = val
		if solved := iterSudoku(board, openSq, openVals, iter + 1); solved {
			return true
		}
		(*board)[xys.y][xys.x] = '.' // mark square as empty
		// Failed to find solution; make val available
		//    in that row/col/sect
		openVals.xval[xys.x][val]=true
		openVals.yval[xys.y][val]=true
		openVals.sval[xys.s][val]=true
	}
	return false
}

type x_y_s struct {
	x int  // col
	y int  // row
	s int  // section
}
var x_y_s_vals = genXYS()

// x, y, sect values for each square on board
func genXYS() []x_y_s {
	x_y_s_vals := make([]x_y_s, nBoard)
	for i:=0; i<nBoard; i++ {
		x := i % nXY
		y := i / nXY
		// Board is divided into 9 3x3 sections
		sect := (x / nSect) + nSect * (y / nSect)
		x_y_s_vals[i] = x_y_s{x, y, sect}
	}
	return x_y_s_vals
}

// Values in next open square to iterate over
func findOpenSquares(board *[][]byte) []int {
	openSq := []int{}
	for i, x_y := range x_y_s_vals {
		if (*board)[x_y.y][x_y.x] == '.' {
			openSq = append(openSq, i)
		}
	}
	return openSq
}

// Keep track of avail choices in each row, col, sect
type openVal struct {
	xval []map[byte]bool // values free in x (rows)
	yval []map[byte]bool // values free in y (cols)
	sval []map[byte]bool // values free in 3x3 sections
}	

// Find open row, col, section values for each empty square
func findOpenVals(board [][]byte) (openVal) {
	// Initialize list with all values 1-8 avail
	open := openVal{}
	open.xval = make([]map[byte]bool, nXY)
	open.yval = make([]map[byte]bool, nXY)
	open.sval = make([]map[byte]bool, nXY)
	for i := 0; i<nXY; i++ {
		open.xval[i] = map[byte]bool{}
		open.yval[i] = map[byte]bool{}
		open.sval[i] = map[byte]bool{}
		for j := 0; j < nXY; j++ {
			val := '1' + byte(j)
			open.xval[i][val] = true
			open.yval[i][val] = true
			open.sval[i][val] = true
		}
	}
	// Mark every specified val as unavail in that row, col, sec
	for _, xys := range x_y_s_vals {
		val := board[xys.y][xys.x]
		delete(open.xval[xys.x], val)
		delete(open.yval[xys.y], val)
		delete(open.sval[xys.s], val)
	}
	return open
}

// Valid values - intesection of valid row, col, grid
func findValidVals(iSquare int, open *openVal) []byte {
	intersect := []byte{}
	xys := x_y_s_vals[iSquare]
	for xval := range open.xval[xys.x] {
		_, isYval := open.yval[xys.y][xval]
		_, isSval := open.sval[xys.s][xval]
		// If x, y, s all valid values
		if isYval && isSval {
			intersect = append(intersect, xval)
		}
	}
	return intersect
}	

// Print board configuration ('.' for empty)
func printBoard(board [][]byte) {
	// fmt.Println(string(board[0][1]))
	fmt.Println()
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
	fmt.Print("\n")
}

// Print list of byte values from map
func printOpenVals(open map[byte]bool) {
	for val := range open {
		fmt.Print(string(val), " ")
	}
	fmt.Println()
}

func testBoard() [][]byte {
	test := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	return [][]byte(test)
}

