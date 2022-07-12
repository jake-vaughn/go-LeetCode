package main

import "fmt"

func isPalindrome(x int) bool {
	x1 := 0

	if x < 0 {
		return false
	}

	for i := 10; i <= x*10; i *= 10 {
		x1 *= 10
		x1 += (x % i) / (i / 10)
		fmt.Println((x % i) / (i / 10))
	}

	if x == x1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(121))
}
