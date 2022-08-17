package main

import (
	"fmt"
	"sort"
)

//This is much faster down to 53 ms from 3000 but it is messy
func threeSum(nums []int) [][]int {
	ret := [][]int{}

	uniqueMap := make(map[int]int)
	uniqueSet := make([]int, 0, len(nums)/2)
	for _, n := range nums {
		count := uniqueMap[n]
		if count < 2 || (n == 0 && count < 3) {
			uniqueMap[n]++
			uniqueSet = append(uniqueSet, n)
		}
	}

	sort.Slice(uniqueSet, func(i, j int) bool {
		return uniqueSet[i] < uniqueSet[j]
	})

	fmt.Printf("uniqueSet: %v\n", uniqueSet)

	for i, a := range uniqueSet {
		if i > 0 && a == uniqueSet[i-1] {
			continue
		}
		zeroFlag := false
		l, r := i+1, len(uniqueSet)-1
		for l < r {
			if uniqueSet[l] == uniqueSet[l-1] && l-1 != i {
				if uniqueSet[l] == 0 && !zeroFlag {
					zeroFlag = true
					continue
				}
				l++
				continue
			}
			bcSum := uniqueSet[l] + uniqueSet[r]
			if bcSum == -a {
				ret = append(ret, []int{a, uniqueSet[l], uniqueSet[r]})
				l++
			} else if bcSum < -a {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1, -14, 0, 1, 2, -1, -4, -14, 28}))
}
