# if two strings are anagrams
package main

import (
       "fmt"
)

func main() {
     fmt.Println("adam", "ada",anagram("adam", "ada"))
     fmt.Println("adam", "adam", anagram("adam", "adam"))
     fmt.Println("adam", "mada", anagram("adam", "mada"))
     //fmt.Println(anagram("ææè", "èææ"))
}

func anagram(str1 string, str2 string) bool {
     len1 := len(str1)
     if len1 != len(str2) {
     	return false}
	for pos := 0; pos < len1; pos++ {
	    char1 := str1[pos]
	    char2 := str2[len1 - pos - 1]
	    //fmt.Println(char1, char2)
	    if char1 != char2 {
	     	return false}
}}