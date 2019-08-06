/** 69. Sqrt(x)
Compute and return the square root of x,
  where x is guaranteed to be a non-negative integer.h
e decimal digits are truncated and only the integer part of the result is returned.
**/

package main

import (
	"fmt"
	// "math"
)

func main() {
	fmt.Println(mySqrt(15)) // 3
	fmt.Println(mySqrt(16)) // 4
}

func mySqrt(x int) int {
	// for i := 1; i < 20; i++ {
	//   fmt.Println(i, sqrGuess(i), float64(sqrGuess(i))/math.Sqrt(float64(i)))}
	xGuess := sqrApprox(x)
	return sqr_Newton(x, xGuess)
}

// Find approx sqrt (to within ~sqrt(2)
func sqrApprox(x int) int {
	x = x / 2
	xGuess := 1
	for ; x > 0; x /= 4 {
		xGuess *= 2
	}
	return xGuess
}

// Newton's method (requires an initial guess)
func sqr_Newton(x, xGuess int) (sqr int) {
	for {
		sqrNext := x / xGuess
		if abs(sqrNext-xGuess) < 2 {
			return min(sqrNext, xGuess)
		}

		fmt.Println(xGuess, sqrNext)
		// Average of the two numbers should be better
		// sqrt estimate than either of them
		xGuess = (xGuess + sqrNext) / 2
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
