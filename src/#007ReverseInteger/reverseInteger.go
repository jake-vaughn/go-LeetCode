package main

import (
	"fmt"
)

func reverseInteger(x int) int {
	x1 := x
	ret := 0

	if x < 0 {
		x1 *= -1
	}

	for i := 10; i <= x1*10; i *= 10 {
		ret *= 10
		ret += (x1 % i) / (i / 10)
		//fmt.Println((x1 % i) / (i / 10))

	}
	if x < 0 {
		ret *= -1
	}
	if ret < -2147483648 || ret > 2147483647 {
		return 0
	}
	return ret
}

func main() {
	x := 7463847412

	fmt.Println(reverseInteger(x))
}
