package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.
*/
func threeSumClosest(nums []int, target int) int {
	//Set return "closest" to first sum
	ret := nums[0] + nums[1] + nums[2]
	distance := math.Abs(float64(target - ret))

	//Sorts nums smallest to largest
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i, a := range nums {
		if i > 0 && a == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l] == nums[l-1] && l-1 != i {
				l++
				continue
			}
			abcSum := a + nums[l] + nums[r]
			if abcSum == target {
				return target
			}
			if math.Abs(float64(target-abcSum)) < distance {
				ret = abcSum
				distance = math.Abs(float64(target - abcSum))
			}
			if abcSum > target {
				r--
			} else {
				l++
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSumClosest([]int{3, 0, 0, 0}, 1))
}
