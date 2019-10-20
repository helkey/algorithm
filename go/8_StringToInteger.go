/* 8. String to Integer (atoi)
  Faster than 52% of Go submissions
*/

package main

import "fmt"

func main() {
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("  -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("91283472332"))
}
	
func myAtoi(str string) int {
	const maxInt = (1 << 31) - 1
	const minInt = (1 << 31) // abs(min integer)
	// check for valid number
	foundNum := false
	var i int
	for i=0; i<len(str); i++ {
		if (str[i] == '-') || (str[i] == '+') || numeric(str[i]) {
			foundNum = true
			break
		}
		if str[i] != ' ' {
			return 0
		}
	}
	if !foundNum {
		return 0
	}
	sign := 1
	// check for minus sign
	if (str[i] == '-') || (str[i] == '+') {
		if str[i] == '-' {
			sign = -1
		}
		i += 1
		if i == len(str) || !numeric(str[i]) {
			return 0 // no numeric chars
		}
	}

	sum := int(str[i] - '0')
	// take numeric digits
	for j:=i+1; j<len(str); j++ {
		if !numeric(str[j]) {
			return sign * sum
		}
		sum = sum * 10 + int(str[j] - '0')
		// check for overflow
		if sign == 1 &&  sum > maxInt {
			return maxInt
		}
		if sign != 1 && sum > minInt {
			return -minInt
		}
	}
	return sign * sum
}

func numeric(b byte) bool {
	return (b >= '0') && (b <= '9')
}
