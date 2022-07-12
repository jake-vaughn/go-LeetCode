package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	sorted := []int{}

	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i < len(nums1) {
			if j < len(nums2) {
				if nums1[i] < nums2[j] {
					sorted = append(sorted, nums1[i])
					i++
				} else {
					sorted = append(sorted, nums2[j])
					j++
				}

			} else {
				sorted = append(sorted, nums1[i])
				i++
			}
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	sLen := len(sorted)
	if sLen == 0 {
		return 0
	}
	if sLen%2 == 0 {
		a := int(sLen / 2)
		b := a - 1
		add := float64(sorted[a] + sorted[b])
		return add / 2
	} else {
		return float64(sorted[sLen/2])
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
