// https://leetcode.com/problems/number-of-islands/discuss/413292/Python-BFS

/* 200. Number of Islands
    https://leetcode.com/problems/number-of-islands/
  Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water 
  and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid 
  are all surrounded by water.

Faster than 100% of Go submissions */

package main

import "fmt"

/* Example 1: (output: 1)
Input:
11110
11010
11000
00000 */

var grid = toByte([]string{"11000", "11000", "00100", "00011"}) // 3

var gridL = [][]byte{
	[]byte{'1','0','0','1','1','1','0','1','1','0','0','0','0','0','0','0','0','0','0','0'},
        []byte{'1','0','0','1','1','0','0','1','0','0','0','1','0','1','0','1','0','0','1','0'},
	[]byte{'0','0','0','1','1','1','1','0','1','0','1','1','0','0','0','0','1','0','1','0'},
	[]byte{'0','0','0','1','1','0','0','1','0','0','0','1','1','1','0','0','1','0','0','1'},
	[]byte{'0','0','0','0','0','0','0','1','1','1','0','0','0','0','0','0','0','0','0','0'},
	[]byte{'1','0','0','0','0','1','0','1','0','1','1','0','0','0','0','0','0','1','0','1'},
	[]byte{'0','0','0','1','0','0','0','1','0','1','0','1','0','1','0','1','0','1','0','1'},
	[]byte{'0','0','0','1','0','1','0','0','1','1','0','1','0','1','1','0','1','1','1','0'}, 
	[]byte{'0','0','0','0','1','0','0','1','1','0','0','0','0','1','0','0','0','1','0','1'},
	[]byte{'0','0','1','0','0','1','0','0','0','0','0','1','0','0','1','0','0','0','1','0'},
	[]byte{'1','0','0','1','0','0','0','0','0','0','0','1','0','0','1','0','1','0','1','0'},
	[]byte{'0','1','0','0','0','1','0','1','0','1','1','0','1','1','1','0','1','1','0','0'},
	[]byte{'1','1','0','1','0','0','0','0','1','0','0','0','0','0','0','1','0','0','0','1'},
	[]byte{'0','1','0','0','1','1','1','0','0','0','1','1','1','1','1','0','1','0','0','0'},
	[]byte{'0','0','1','1','1','0','0','0','1','1','0','0','0','1','0','1','0','0','0','0'},
	[]byte{'1','0','0','1','0','1','0','0','0','0','1','0','0','0','1','0','1','0','1','1'},
	[]byte{'1','0','1','0','0','0','0','0','0','1','0','0','0','1','0','1','0','0','0','0'},
	[]byte{'0','1','1','0','0','0','1','1','1','0','1','0','1','0','1','1','1','1','0','0'},
	[]byte{'0','1','0','0','0','0','1','1','0','0','1','0','1','0','0','1','0','0','1','1'},
	[]byte{'0','0','0','0','0','0','1','1','1','1','0','1','0','0','0','1','1','0','0','0'}}

func toByte(strs []string) [][]byte {
	bArr := make([][]byte, len(strs))
	for i:=0; i<len(strs); i++ {
		bArr[i] = []byte(strs[i])
	}
	return bArr
}

func printArr(bArr [][]byte) {
	for _, bytes := range bArr {
		// fmt.Println(bytes)
		fmt.Println(string(bytes))
	}
}

func main() {
	// fmt.Println(grid)
	fmt.Println(numIslands(grid))
	printArr(gridL)
}



func numIslands(grid [][]byte) int {
	nY := len(grid)
	if (nY == 0) || (len(grid[0]) == 0) {
		return 0
	}
	nX := len(grid[0])
	color := byte(2) // start after color=1 -> presence of land
	for x:=0; x<nX; x++ {
		for y:=0; y<nY; y++ {
			if grid[y][x] == '1' {
				colorIsland(grid, color, nX, nY, x, y)
				color += 1 
			}
		}
	}
	return int(color - 2) // color=1 not used
}

// Change color of contiguous island to 'color'
func colorIsland(grid [][]byte, color byte, nX, nY, x, y int) {
	if x < 0 || x >= nX  || y < 0 || y >= nY || grid[y][x] != '1' {
		return  // no island found
	}
	grid[y][x] = '0' // + byte(color) // // useful for debugging
	// check adjacent grid squares
	colorIsland(grid, color, nX, nY, x - 1, y)
	colorIsland(grid, color, nX, nY, x + 1, y)
	colorIsland(grid, color, nX, nY, x,     y - 1)
	colorIsland(grid, color, nX, nY, x,     y + 1)
}



