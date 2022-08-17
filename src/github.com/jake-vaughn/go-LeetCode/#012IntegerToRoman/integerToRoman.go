package main

import "fmt"

func stringNTimes(s string, n int) string {
	ret := ""
	for i := 1; i <= n; i++ {
		ret += s
	}
	return ret
}

func intToRoman(num int) string {
	ret := ""
	one, five, ten := "", "", ""

	for i := 1; num > 0; i++ {
		switch i {
		case 1:
			one = "I"
			five = "V"
			ten = "X"
		case 2:
			one = "X"
			five = "L"
			ten = "C"
		case 3:
			one = "C"
			five = "D"
			ten = "M"
		case 4:
			one = "M"
		}
		n := num % 10
		switch {
		case n < 4:
			ret = stringNTimes(one, n) + ret
		case n == 4:
			ret = one + five + ret
		case n == 9:
			ret = one + ten + ret
		case n >= 5:
			ret = five + stringNTimes(one, n-5) + ret
		}
		num /= 10
	}
	return ret
}

func main() {
	fmt.Println(intToRoman(28))
}
