package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, num := range numbers {
		if idx, ok := m[target-num]; ok {
			return []int{idx + 1, i + 1}
		}
		m[num] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
