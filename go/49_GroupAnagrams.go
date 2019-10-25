/* 49. Group Anagrams
    Given an array of strings, group anagrams together.

(faster than 84%) */

package main
import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

//struct strArr {
//	[]string
// 

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	// Find anagrams by sorting string, put in hashmap
		for _, str := range strs {
		keyCh := strings.Split(str, "")
		sort.Strings(keyCh)
		key := strings.Join(keyCh, "")
		if _, ok := hash[key]; !ok {
			hash[key] = []string{str}
		} else {
			hash[key] = append(hash[key], str)
		}
	}
	
	out := [][]string{}
	for _, strArr := range(hash) {
		out = append(out, strArr)
	}
	return out
}
