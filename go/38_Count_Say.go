/** 38. Count and Say

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Runtime: 4 ms, faster than 64.97% of Go online submissions for Count and Say.
Memory Usage: 6.6 MB, less than 100.00% of Go online submissions for Count and Say.

**/

package main

import "fmt"

func main() {
	fmt.Println(countAndSay(31))
}

func countAndSay(n int) string {
	str := "1"
	for l:=2; l<=n; l++ {
		new := ""
		lenStr := len(str)
		for i:=0; i<lenStr; i++ {
			cNext := str[i]
			count := 1
			for j:=i+1; j<lenStr; j++ {
				if str[j] != cNext {
					break
				}
				count++
			}
			new += string('0' + count) + string(cNext)
			i += count - 1
		}
		str = new
	}
	return str
}
