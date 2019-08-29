package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(fizzBuzzStr(5)) // 'buzz'

	for i:=1; i<=20; i++ {
		fmt.Println(fizzBuzzStr(i))
	}
}


// Generate fizzbuzz string
func fizzBuzzStr(i int) string {
	s := ""
	if (i%3) == 0 {
		s = "fizz"
	}
	if (i%5) == 0 {
		s += "buzz"
	}
	if s == "" {
		s = strconv.Itoa(i)
	}
	return s
}
