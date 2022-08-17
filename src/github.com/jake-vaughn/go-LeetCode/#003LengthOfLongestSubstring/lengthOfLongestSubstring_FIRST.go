package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	fmt.Println(s)
	max := 0
	m := make(map[string]int)

	for i, c := range s {
		c := string(c)

		if _, ok := m[c]; ok {
			max = i
			spl := strings.SplitN(s, string(s[m[c]]), 2)
			recmin := lengthOfLongestSubstring(spl[1])
			fmt.Println(max, recmin)

			if recmin > max {
				return recmin
			}
			return max
		}
		m[c] = i
		fmt.Println(m)

	}
	return len(s)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
