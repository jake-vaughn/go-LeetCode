package main

import (
	"fmt"
)

//returns n to the power m
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

//Interface for counter
type counter interface {
	increment()
}

//Structure for a arbitray length and base counter
//M...M where number of M's = length
// and where M is a digit 0-M where M = base
type lenNBaseMCounter struct {
	vars   []int
	length int
	base   int
}

//Icrements counter by one overflowing to 0 if counter is full
func (c lenNBaseMCounter) increment() {
	if c.length <= 0 || c.base == 0 {
		return
	}
	if c.vars == nil {
		c.vars = make([]int, c.length)
	}
	//increments counter then checks for overflow
	c.vars[0]++
	for i := 0; i < c.length; i++ {
		if c.vars[i] == c.base {
			if i+1 < c.length {
				c.vars[i+1]++
				c.vars[i] = 0
			} else {
				c.vars = make([]int, c.length)
			}
		} else {
			return
		}
	}
}

func letterCombinations(digits string) []string {
	ret := []string{}
	digitMap := map[rune][4]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	dRunes := []rune(digits)
	//initializes counter with the length of digits
	// and a base of 4
	count := lenNBaseMCounter{
		vars:   make([]int, len(digits)),
		length: len(digits),
		base:   4,
	}

	//Loops for every combination of letters using the counter
	// to index each letter
	for j := 0; j < IntPow(count.base, count.length); j++ {
		str := ""
		for i := 0; i < count.length; i++ {
			idx := count.vars[count.length-1-i]
			letter := digitMap[dRunes[i]][idx]
			if letter == "" {
				str = ""
				break
			}
			str += letter
		}
		if str != "" {
			ret = append(ret, str)
		}
		count.increment()
	}
	return ret
}

func main() {
	fmt.Println(letterCombinations("563"))
}
