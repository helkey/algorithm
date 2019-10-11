package main

import (
       "fmt"
)

func main() {
     test_replace()
}

func replace_space(str_in string) string {
     i_start := 0
     str_out := ""
     var substr string
     for i:=0; i<len(str_in); i++ {
     if str_in[i] == ' ' {
     	if i_start == i {
	   substr = ""
	} else {
	  substr = str_in[i_start: i]
 	}
	str_out = str_out + substr + "%20"
	i_start = i + 1
	}
    }
    if i_start <= len(str_in) - 1 {
	str_out += str_in[i_start: len(str_in)]
	}
	return str_out
}

func test_replace() {
     fmt.Println(replace_space("test replace"))
     fmt.Println(replace_space("test  replace"))
     fmt.Println(replace_space(" test replace "))
}


