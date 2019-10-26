/* 29. Divide Two Integers

(faster than 100% of Go submissions */

package main
import "fmt"

func main() {
	fmt.Println(divide(2559, 10))        // 255
	fmt.Println(divide(10, 3))           // 3
	fmt.Println(divide(-2147483648, -1)) // Expected: 2147483647
}

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return (1<<31) - 1
	}
	pos := true
	if ((dividend < 0) && (divisor > 0)) || ((dividend > 0) && (divisor < 0)) {
		pos = false
	}
	dvd := abs(dividend)
	dvr := abs(divisor)
	bit := 1
	dvr1 := dvr
	// Find highest bit of ratio d
	for ; dvr1 * 2 <= dvd; dvr1 = dvr1 << 1 {
		bit = bit << 1
		// fmt.Println(dvr1, bit)
	}
	div := 0
	for ; (bit > 0) && (dvd != 0); bit = bit >> 1 {
		if dvr1 <= dvd {
			dvd = dvd - dvr1
			div = div + bit
		}
		dvr1 = dvr1 >> 1
		// fmt.Println(dvd, dvr1, div, bit)
   	}
	// fmt.Println(bit, pos)
	maxDiv := 1 << 31 - 1
	minDiv := -(1 << 31)
	if pos {
		return min(div, maxDiv)
	}
	return max(-div, minDiv)
}

func abs(x int) int {
	if x>=0 {
		return x
	}
	return -x
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
