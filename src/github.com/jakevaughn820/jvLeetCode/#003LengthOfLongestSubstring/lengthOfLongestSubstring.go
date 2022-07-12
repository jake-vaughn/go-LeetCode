package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//Window Question
func lengthOfLongestSubstring(s string) int {
	res, i := 0, 0
	m := make(map[rune]bool)

	for _, b := range s {
		//While current rune is in Map remove rune at j in s
		for m[b] {
			delete(m, rune(s[i]))
			i++
		}
		//Add rune b to map
		m[b] = true
		//len(m) is current HashMap window size
		res = max(len(m), res)
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
