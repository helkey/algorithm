/* 367. Valid Perfect Square
  https://leetcode.com/problems/valid-perfect-square/

Faster than 100% of Go submissions
*/

package main
import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16)) // assert true
	fmt.Println(isPerfectSquare(14)) // assert false
    
}

// Binary search for integer square root O(log(N))
func isPerfectSquare(num int) bool {
	if num < 9 {
		return (num == 1 || num == 4)
	}
	// binary search
	r1, r2 := 2, num/2
	for ; r2 > r1 + 1; {
		rNew := (r1 + r2) / 2
		sqr := rNew * rNew
		// fmt.Println(r1, r2, rNew, sqr)
		if sqr == num {
			return true
		}
		if sqr < num {
			r1 = rNew
		} else {
			r2 = rNew
		}
	}
	return false
}
