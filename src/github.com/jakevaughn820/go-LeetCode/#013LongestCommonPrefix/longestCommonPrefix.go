package main

import "fmt"

/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string ""
*/
func longestCommonPrefix(strs []string) string {
	ret := ""
	minLen := 200
	//Loop to find smallest string in strs
	for _, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	//Loop to iterate through characters to compare
	for i := 0; i < minLen; i++ {
		//Loop to iterate through strs
		for j := 0; j < len(strs)-1; j++ {
			//If the bytes from the string are not equal return
			if strs[j][i] != strs[j+1][i] {
				return ret
			}
		}
		//All characters match add to string
		ret += string(strs[0][i])
	}

	return ret
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"low", "lomba", "loon"}))
}
