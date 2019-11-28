# Fast Sudoku Solver

Sudoku puzzles are a popular pastime, and a good opportunity to illustrate recursion and backtracking.

An [example of a sudoku puzzle](https://leetcode.com/problems/sudoku-solver/) is shown below:

![Sudoku problem](/250px-Sudoku-by-L2G-20050714.svg.png "Sudoku problem")

A Sudoku solver is a classic example of recursive backtracking. This algorithm makes a list of possible choices at each square, 
  tries one of the possible choices, then passes the new set with additional choice to a recursive solver. 
  If all of the possible choices fail to produce a solution, it backtracks up the stack to an earlier selection and tries another choice.

Solving a hard Sukoku problem can require searching a large possible solution set. In order to make this algorithm run quickly, 
  we want to pre-calculate as much as possible, so that each step runs quickly.

The most important optimization is to minimize the effort in calculating the available choices for each new square. 
  This involves finding the possible choices not already used in the same row, column, and section. 
  Here the program starts by finding the available values in each row, column and section. 
  The union of these values will be the available choices for each empty square:
```
     // Find open row, col, section valuesfor each empty square
     func findOpenVals(board [][]byte) (openVal) {
          // Initialize list with all values 1-8 avail
          open := openVal{}
          open.xval = make([]map[byte]bool, nXY)
		...
```
As a value is selected for each square, it is removed from the list of available row, col, and section values:
```
     // Selected val no longer available in the row/col/sect
     delete(openVals.xval[xys.x], val)
     delete(openVals.yval[xys.y], val)
     delete(openVals.sval[xys.s], val)
```
If this selected value fails to produce a solution, the backtracking algorithm removes this value from the square, 
  and adds it back in to the available values for squares in the same row, column, or section:
```
     // Failed to find solution; make val available
     //    in that row/col/sect
     openVals.xval[xys.x][val]=true
     openVals.yval[xys.y][val]=true
     openVals.sval[xys.s][val]=true
```
An additional efficiency improvement to minimize calculation at each iteration is generate a list of all of the empty squares in a list, 
  so that at each iteration the program is not looking for the next open square:
```
     func findOpenSquares(board *[][]byte) []int {
          openSq := []int{}
          for i, x_y := range x_y_s_vals {
               if (*board)[x_y.y][x_y.x] == '.' {
                    openSq = append(openSq, i)
               }
          }
          return openSq
     }
```

Sudoku problems are expected to be designed to have only a single solution.
This program returns the first solution:

![Sudoku solution](/250px-Sudoku-by-L2G-20050714_solution.svg.png "Sudoku solution")

but could be trivially extended to check if there is more than one possible solution.

