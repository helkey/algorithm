package main
import "fmt"

func main() {
	fmt.Println(convert2str("12345678"))
}

func convert2str(s string) int {
	dec := 0
	for i:=0; i<len(s); i++ {
		c := s[i] - '0'
		dec = 10 * dec + int(c)
		fmt.Println(c)
	}
	return dec
}
		
