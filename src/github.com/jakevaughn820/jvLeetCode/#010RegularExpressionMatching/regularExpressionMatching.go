package main

import "fmt"

func isMatch(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	wCount, reqCount := 0, 0
	//wildUsed := false

	wilds := make([]byte, 0, lenP/2)
	required := make([]byte, 0, lenP)

	for j := lenP - 1; j >= 0; j-- {
		if string(p[j]) == "*" {
			j--
			wilds = append(wilds, p[j])
		} else {
			required = append(required, p[j])
		}
	}

	fmt.Println(lenS, wilds, required)
	for i := lenS - 1; i >= 0; i-- {
		if required[reqCount] == byte(97) || required[reqCount] == s[i] {
			reqCount++
		} else {
			if wilds[wCount] == byte(97) || wilds[wCount] == s[i] {

			}
		}
	}

	return false
}

func main() {
	fmt.Println(isMatch(
		"aaaaa",
		"aa*"))
}
