package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxArea(height []int) int {
	width := len(height)
	maxWater := 0

	l := 0
	for r := width - 1; l < r; r-- {
		rHgt := height[r]
		if rHgt*(r-l) > maxWater {
			for l < r && height[l]*(r-l) <= maxWater {
				l++
			}
			waterArea := min(rHgt, height[l]) * (r - l)
			if waterArea > maxWater {
				maxWater = waterArea
				r++
			}
		}
	}
	return maxWater
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
