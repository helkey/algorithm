package main

import (
	"fmt"
)

type list []interface{}

func main() {
	// [[1,2,[3]],4] -> [1,2,3,4]
	lst := list{list{1, 2, list{3}}, 4}
	fmt.Println(flatten(lst))
}

func flatten(lst []interface{}) list {
	flat := make(list, 0)
	for _, v := range lst {
		switch v.(type) {
		case list: // type []interface{}
			flat = append(flat, flatten(v.(list))...)
		default: // types int, float64, string, bool...
			flat = append(flat, v)
		}
	}
	return flat
}
