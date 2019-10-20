/* 6. ZigZag Conversion
  The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: r legibility)
  P   A   H   N
  A P L S I I G
  Y   I   R
And then read line by line.
  (Faster than 96% of Go solutions) */

package main

import "fmt"

func main() {
	convert("PAYPALISHIRING", 4)
}

func convert(s string, numRows int) string {
	if (len(s) <= 1) || (numRows <= 1) {
		return s
	}
	// (x,y) grid position array
	sArr := make([][]byte, numRows)
	rt := false
	y := 0
	for i:=0; i<len(s); i++ {
		sArr[y] = append(sArr[y], s[i])
		if y == 0 {
			rt = false
		} else if y == numRows - 1 {
			rt = true
		}
		if !rt {
			y += 1
		} else {
			y -= 1
		}
	}
	sOut := sArr[0]
	for r:=1; r<numRows; r++ {
		sOut = append(sOut, sArr[r]...)
	}
	return string(sOut)
}
