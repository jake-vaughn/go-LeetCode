package main

import (
	"fmt"
	"sort"
)

//returns the absolute value
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.
*/
func threeSumClosest(nums []int, target int) int {
	ret := nums[0] + nums[1] + nums[2]
	diff := abs(target - ret)
	arrLen := len(nums) - 1

	//Sorts nums smallest to largest
	sort.Ints(nums)

	for i, a := range nums {
		l, r := i+1, arrLen
		for l < r {
			abcSum := a + nums[l] + nums[r]
			if abs(target-abcSum) < diff {
				ret = abcSum
				diff = abs(target - abcSum)
			}
			if abcSum > target {
				r--
			} else if abcSum < target {
				l++
			} else {
				return target
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSumClosest([]int{3, 0, 0, 0}, 1))
}
