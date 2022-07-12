package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums1 := []int{3, 7, 2, 5, 8}

	fmt.Println(twoSum(nums1, 9))

}
