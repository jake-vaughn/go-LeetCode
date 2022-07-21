package main

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/
func threeSum(nums []int) [][]int {
	ret := [][]int{}

	//Sorts nums smallest to biggest
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	//Ranges over nums to find triplets
	for i, a := range nums {
		//if a is = to last num skip it
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			//Skips checking nums[l] if it is the same as last
			if nums[l] == nums[l-1] && l-1 != i {
				l++
				continue
			}
			bcSum := nums[l] + nums[r]
			if bcSum == -a {
				ret = append(ret, []int{a, nums[l], nums[r]})
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
	fmt.Println(threeSum([]int{-1, -14, 0, 1, 0, 0, 0, 2, -1, -4, -14, 28}))
}
