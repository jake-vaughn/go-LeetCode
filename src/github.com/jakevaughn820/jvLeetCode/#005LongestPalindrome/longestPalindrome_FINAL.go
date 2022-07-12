package main

import (
	"fmt"
)

//Index Expansion
func longestPalindrome(s string) string {
	sLen := len(s)
	paliLen := -1
	lRet, rRet := 0, 0
	for lMid, rMid := 0, 0; rMid < sLen; {
		for i := 0; (lMid-i) >= 0 && (rMid+i) < sLen && s[lMid-i] == s[rMid+i]; i++ {
			fmt.Println(lMid, rMid)
			fmt.Println(string(s[lMid-i]), string(s[rMid+i]))
			if (rMid+i)-(lMid-i) > paliLen {
				paliLen = (rMid + i) - (lMid - i)
				lRet = lMid - i
				rRet = rMid + i + 1
				fmt.Println("ret is now", s[lRet:rRet])
			}
		}

		if lMid == rMid {
			rMid++
		} else {
			lMid++
		}
	}
	return s[lRet:rRet]
}

func main() {
	s := "abccba"
	fmt.Println(longestPalindrome(s))
}
