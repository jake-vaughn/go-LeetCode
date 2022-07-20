package main

import "fmt"

func romanToInt(s string) int {
	ret := 0

	flagI, flagX, flagC := false, false, false

	for _, r := range s {
		switch r {
		case 'I':
			ret += 1
			flagI = true
		case 'V':
			if flagI {
				ret -= 2
			}
			ret += 5
		case 'X':
			if flagI {
				ret -= 2
				flagI = false
			}
			flagX = true
			ret += 10
		case 'L':
			if flagX {
				ret -= 20
			}
			ret += 50
		case 'C':
			if flagX {
				ret -= 20
				flagX = false
			}
			flagC = true
			ret += 100
		case 'D':
			if flagC {
				ret -= 200
			}
			ret += 500
		case 'M':
			if flagC {
				ret -= 200
				flagC = false
			}
			ret += 1000
		}
	}
	return ret
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
