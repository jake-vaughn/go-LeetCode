package main

import (
	"fmt"
	"sort"
)

/* Given an array nums of n integers, return an array of
all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target

You may return the answer in any order. */
func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	nLen := len(nums)
	sort.Ints(nums)
	for i := 0; i < nLen-3; i++ {
		a := nums[i]
		if i != 0 && a == nums[i-1] {
			continue
		}
		for j := i + 1; j < nLen-2; j++ {
			b := nums[j]
			if j != i+1 && b == nums[j-1] {
				continue
			}
			l, r := j+1, nLen-1
			for l < r {
				c, d := nums[l], nums[r]
				if l != j+1 && c == nums[l-1] {
					l++
					continue
				}
				abcdSum := a + b + c + d
				if abcdSum == target {
					ret = append(ret, []int{a, b, c, d})
					l++
				}
				if abcdSum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2, -1}, 0))
}
