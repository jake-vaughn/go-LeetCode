package main

import "fmt"

func isMatch(s string, p string) bool {
	//wilds stores "*" chars until used (max len is len(p)/2)
	wilds := make([]string, len(p)/2)
	//wCur tracks current wild in wilds
	//wIdx tracks index of next saved wild in wilds
	wCur, wIdx := 0, 0
	//Backwards iterators for s and p
	i, j := len(s)-1, len(p)-1

	//iterate over s and p until one ends
	for i >= 0 && j >= 0 {
		chrS, chrP := string(s[i]), string(p[j])
		//If current char in p is a "*" save it in wilds
		//then iterate p keeping place in s the same
		if string(p[j]) == "*" {
			j--
			chrP = string(p[j])
			wilds[wIdx] = chrP
			wIdx++
			j--
			continue
		}
		//If current char p reg matches current s
		if chrP == "." || chrP == chrS {
			//If there are no wild chars in wilds iterate and continue
			if wCur == wIdx {
				i--
				j--
				continue
			}
			//If there are wild chars continue RECURSIVELY with truncated s and p
			//If false is returned wildcards will be used to try and continue
			//If true is returned return return true
			if isMatch(s[:i], p[:j]) {
				return true
			}
		}
		//While there are wild chars
		for wCur < wIdx {
			//If the current wild char matches s iterate break
			if string(wilds[wCur]) == "." || wilds[wCur] == chrS {
				break
			}
			//iterate wilds
			wCur++
		}
		//If no wild chars return false
		if wCur == wIdx {
			return false
		}
		//Iterate s
		i--
	}

	//Final Check if S has not been iterated through all the way.
	for i >= 0 {
		//If no "*" characters left return false
		if wCur == wIdx {
			return false
		}
		//If the "*" char matches iterate S if not iterate through "*" chars
		if string(wilds[wCur]) == "." || wilds[wCur] == string(s[i]) {
			i--
		} else {
			wCur++
		}
	}
	//Final Check if P has not been iterated through all the way.
	//If j contains anything other than wildcard "*" chars return false
	for j >= 0 {
		if string(p[j]) != "*" {
			return false
		}
		j -= 2
	}
	//everything passes return true
	return true
}

func main() {
	fmt.Println(isMatch(
		"aab",
		"c*a*b"))
}
