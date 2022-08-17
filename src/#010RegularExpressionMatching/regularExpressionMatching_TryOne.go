package main

import "fmt"

func isMatch(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	i, j := 0, 0

	for i, j = 0, 0; i < lenS && j < lenP; i, j = i+1, j+1 {

		if j+1 < lenP && string(p[j+1]) == "*" {
			fmt.Println(string(s[i]), string(p[j])+"*")

			if s[i] != p[j] && string(p[j]) != "." {
				j++
				i--
				continue
			}
			fmt.Println("sending ", s[i+1:], " and ", p[j+2:])
			m := isMatch(s[i:], p[j+2:])
			fmt.Println(m)
			if m {
				return true
			}

			j--
			continue
		}
		fmt.Println(string(s[i]), string(p[j]))

		if s[i] != p[j] && string(p[j]) != "." {
			return false
		}
	}
	//If both i and j reached the end of both strings s and p return true
	if i == lenS && j == lenP {
		return true
	}
	//If i reached length of s and string p ended in * return true
	if i == lenS && j == lenP-2 && string(p[j+1]) == "*" {
		return true
	}
	return false
}

func main() {
	fmt.Println(isMatch(
		"aabcbcbcaccbcaabc",
		".*a*aa*.*b*.c*.*a*"))
}
