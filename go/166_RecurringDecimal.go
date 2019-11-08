/* 166. Fraction to Recurring Decimal 
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format. If the fractional part is repeating, enclose the repeating part in parentheses.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(1, 2)) // Output: "0.5"}
	fmt.Println(fractionToDecimal(2, 1)) // Output: "2"}
	fmt.Println(fractionToDecimal(2, 3)) // Output: "0.(6)"}
	fmt.Println(fractionToDecimal(4, 333)) // Output: "0.(012)"}
	fmt.Println(fractionToDecimal(1, 6)) // Output: "0.1(6)"}
	fmt.Println(fractionToDecimal(1, 17)) // Output: "0.(0588235294117647)"}
	fmt.Println(fractionToDecimal(-50, 8)) // Output: "-6.25"}
	fmt.Println(fractionToDecimal(7, -12)) // Output: "-0.58(3)"}
}

// Decimal must either repeat or terminate
func fractionToDecimal(numerator int, denominator int) string {
	str := ""
	if numerator * denominator < 0 {
		str = "-"
	}
	numer, denom := abs(numerator), abs(denominator)
	intr := numer/denom
	str += strconv.Itoa(intr)
	rem := 10 * (numer % denom)
	// fmt.Println(numer, denom, rem, str)
	if rem == 0 {
		return str
	}
	str += "."
	rems := map[int]int{}
	dig := rem / denom
	// fmt.Println(str, dig, rem)
	for i := len(str); ; i++ {
		// If remainder same, strip back repeated section, put in parens	
		if _, ok := rems[rem]; ok {
			str = str[:rems[rem]] +  "(" + str[rems[rem]:] + ")"
			return str
		}
		rems[rem] = i
		rem = 10 * (rem % denom)
		nextDig := rem / denom
		// fmt.Println(rem, dig, nextDig)
		str += strconv.Itoa(dig)
		if rem == 0 {
			return str
		}
		dig = nextDig
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
