/* 50. Pow(x, n)
Implement pow(x, n), which calculates x raised to the power n (xn).
  https://leetcode.com/problems/powx-n/

Faster than 100% of Go submissions
*/

package main
import "fmt"

func main() {
	fmt.Println(myPow(2.0, 10)) // assert 1024.0
	fmt.Println(myPow(2.1, 3)) // assert 9.261
	fmt.Println(myPow(2.0, -2)) // assert 0.25
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n % 2 == 1 {
		return x * myPow(x, n - 1)
	}
	return myPow(x * x, n/2)
}
