/**  7. Reverse Integer
     Given a 32-bit signed integer, reverse digits of an integer.
**/

package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
	

func reverse(x int) int {
	y:=0
	for x != 0 {
		// fmt.Println(x, y)
		y = y * 10 + x % 10
		x = x / 10
	}
	if y != int(int32(y)) {
		return 0
	}
	return y
}
