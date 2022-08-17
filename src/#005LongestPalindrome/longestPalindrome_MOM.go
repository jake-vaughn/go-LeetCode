package main

import (
	"fmt"
)

func ReadBkwds(newString string) string {
	wordBkwds := ""
	//fmt.Println(newString)
	for i := len(newString) - 1; i >= 0; i-- {
		wordBkwds = wordBkwds + string(newString[i])
	}
	return wordBkwds
}

func longestPalindrome(s string) string {
	word := s
	numChars := 0
	//newPalindrome := ""

	if len(word) == 1 {
		return word
	} else {
		fmt.Println("word length = ", len(word))
		fmt.Println("word is ", word[numChars:len(word)])
		for pLen := len(word); pLen > 0; pLen-- {
			fmt.Println("palindrome length = ", pLen)
			for start := 0; start <= (len(word) - pLen); start++ {
				fmt.Println("start char = ", start)
				//palindrome := true
				numChars = pLen
				fmt.Println("NumChars = ", numChars)
				for i := start; i <= (len(word))/2; i++ {
					//fmt.Println("comparing ", i, " & ", i+numChars-1)
					//fmt.Println("Comparing ", string(word[i]), " & ", string(word[i+numChars-1]))
					fmt.Println("i = ", i)
					if word[i] != word[i+numChars-1] {
						fmt.Println("Can't be a palindrome")
						break
					} else {
						if numChars > 1 {
							numChars = numChars - 2
						} else {
							fmt.Println("Palindrome found:  ", word[start:pLen])
							return word[start:pLen]
						}
					}
				}
			}
		}
	}
	return ("No Palindrome found!")
}

func main() {
	//s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	s := "cbab"
	fmt.Println(longestPalindrome(s))

}
