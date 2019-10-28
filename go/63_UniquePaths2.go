/* 63. Unique Paths II
  Only move either down or right at any point in time. Reach the bottom-right corner of the grid 
  (marked 'Finish' in the diagram).

(faster than 100% of Go submissions) */

package main
import "fmt"

func main() {
	grid := [][]int{[]int{0,0,0}, []int{0,1,0}, []int{0,0,0}}
	fmt.Println(uniquePathsWithObstacles(grid))
}
	
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	nrow := len(obstacleGrid)
	if nrow == 0 {
		return 0
	}
	ncol := len(obstacleGrid[0])
	if ncol == 0 {
		return 0
	}
	for row:=0; row<nrow; row++ {
		for col:=0; col<ncol; col++ {
			if obstacleGrid[row][col] == 1 {
				// obstacle in this square, so no possible paths
				obstacleGrid[row][col] = 0
			} else if row == 0 && col == 0 {
				obstacleGrid[0][0] = 1
			} else {
				paths := 0
				// Sum possible paths from adjacent row and column	
				if row > 0 {
					paths += obstacleGrid[row - 1][col]
				}
				if col > 0 {
					paths += obstacleGrid[row][col - 1]
				}
				obstacleGrid[row][col] = paths
			}
		}
	}
	return obstacleGrid[nrow - 1][ncol - 1]
}
