package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	lMid, rMid := 0, 0
	fmt.Println(s)

	if len(s)%2 == 0 {
		rMid = len(s) / 2
		lMid = rMid - 1
	} else {
		lMid = (len(s)-1)/2 - 1
		rMid = lMid + 2
	}

	for i := 0; (lMid - i) >= 0; i++ {
		l := s[lMid-i]
		r := s[rMid+i]
		if l != r {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	queue := []string{s}
	qMap := make(map[string]bool)

	for queue != nil {
		str := queue[0]
		if isPalindrome(str) {
			return str
		}
		delete(qMap, str)
		queue = queue[1:]
		lStr := str[:len(str)-1]
		rStr := str[1:]
		if !qMap[lStr] {
			qMap[lStr] = true
			queue = append(queue, lStr)
		}
		if !qMap[rStr] {
			qMap[rStr] = true
			queue = append(queue, rStr)
		}
	}
	return ""
}

func main() {
	s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	fmt.Println(longestPalindrome(s))
}
