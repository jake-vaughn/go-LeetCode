package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	res := 0
	neg, flag := false, false
LOOP:
	for _, r := range s {

		if !flag {
			switch r {
			case ' ':
				continue LOOP
			case '-':
				neg = true
				flag = true
				continue LOOP
			case '+':
				flag = true
				continue LOOP
			}
		}
		res *= 10
		switch r {
		case '0':
		case '1':
			res += 1
		case '2':
			res += 2
		case '3':
			res += 3
		case '4':
			res += 4
		case '5':
			res += 5
		case '6':
			res += 6
		case '7':
			res += 7
		case '8':
			res += 8
		case '9':
			res += 9
		default:
			res /= 10
			break LOOP
		}
		flag = true
		if res > math.MaxInt32 {
			if neg {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	if neg {
		res *= -1
	}
	return res
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
