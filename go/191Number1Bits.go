/* 191. Number of 1 Bits
Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).

Faster than 100% of Go submissions
*/ 	

package main
import "fmt"

func main() {
	n := uint32(11)
	fmt.Println(hammingWeight( n )) // 0b1011 = 3
	// fmt.Println(fammingWeight(
	// fmt.Println(fammingWeight(
				
}

func hammingWeight(num uint32) int {
	sum := 0
	for ; num > 0; num = (num >> 1) {
		sum += int(num & 1)
	}
	return sum
}
