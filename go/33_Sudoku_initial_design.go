/** 37. Sudoku Solver
Each of the digits 1-9 must occur exactly once in each row.
	Each of the digits 1-9 must occur exactly once in each column.
	Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
**/

package main

import "fmt"


const nxy = 9 // length of row/cols
const nsect = 3 // len/number of subsections

type openVal struct { // lists of open choices per row, col, seg
	xval []map[byte]bool // values free in x (rows)
	yval []map[byte]bool // values free in y (cols)
	gval []map[byte]bool // values free in 3x3 grids
}	

func main() {
	board := testBoard()
	printBoard(board)
	solveSudoku(board)
	printBoard(board)
}

func solveSudoku(board [][]byte) bool  {
	open := findOpenVals(board)
	solved := iterSudoku(&board, &open, 0)
	if solved { fmt.Println("SOLVED") } else {
		fmt.Println("NOT solved") }
	return solved
}

// Select one value, iterate/backtrack
func iterSudoku(board *[][]byte, open *openVal, iter int) bool {
	x, y, vals, solved := nextVals(board, open)
	// if solved || (iter > 9) {
	if solved {
		return true
	}
	grid := gridNum(x, y)
	// fmt.Printf("level%v  x:%v y:%v g:%v\n", iter, x,  y, grid)
	for _, val := range vals {
		delete(open.xval[x], val)
		delete(open.yval[y], val)
		delete(open.gval[grid], val)
		(*board)[y][x] = val
		if solved := iterSudoku(board, open, iter + 1); solved {
			return true
		}
		// Revert 'val'; failed to find solution
		(*board)[y][x] = '.'
		open.xval[x][val]=true
		open.yval[y][val]=true
		open.gval[grid][val]=true
	}
	return false
}

// Values in next open square to iterate over
func nextVals(board *[][]byte, open *openVal) (int, int, []byte, bool) {
	for y := 0; y<nxy; y++ {
		for x := 0; x< nxy; x++ {
			val := (*board)[y][x]
			if val != '.' {
				continue
			}
			vals := []byte{}
			grid := gridNum(x, y)
			// fmt.Println(x, y, grid)
			// printOpenVals(open.xval[y])
			// printOpenVals(open.yval[y])
			// printOpenVals(open.gval[gridNum(x, y)])
			for xval := range open.xval[x] {
				_, isYval := open.yval[y][xval]
				_, isGval := open.gval[grid][xval]
				if isYval && isGval {
					vals = append(vals, xval)
				}
			}
			return x, y, vals, false
		}
	}
	return 0, 0, []byte{}, true // solution complete
}

func printIterVals(iterVals []byte) {
	for _, v := range iterVals {
		fmt.Print(string(v), " ")
	}
	fmt.Println()
}

func findOpenVals(board [][]byte) (openVal) {
	open := initOpenVals()
	for x := 0; x<nxy; x++ {
		for y := 0; y< nxy; y++ {
			val := board[y][x]
			delete(open.xval[x], val)
			delete(open.yval[y], val)
			grid := gridNum(x, y) // which of 9 grids
			delete(open.gval[grid], val)
		}
	}
	return open
}

func gridNum(ix int, iy int) int {
	sectx := ix / nsect
	secty := iy / nsect
	return nsect * secty + sectx
}
	
func initOpenVals() openVal {
	var open openVal
	open.xval = make([]map[byte]bool, nxy)
	open.yval = make([]map[byte]bool, nxy)
	open.gval = make([]map[byte]bool, nxy)
	for i := 0; i<nxy; i++ {
		open.xval[i] = map[byte]bool{}
		open.yval[i] = map[byte]bool{}
		open.gval[i] = map[byte]bool{}
		for j := 0; j < nxy; j++ {
			val := '1' + byte(j)
			open.xval[i][val] = true
			open.yval[i][val] = true
			open.gval[i][val] = true
		}
	}
	return open
}

// print board configuration ('.' for empty)
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

// print list of byte values from map
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

