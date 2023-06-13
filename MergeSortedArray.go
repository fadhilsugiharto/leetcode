package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for j >= 0 {
		if nums1[i] > nums2[j] && i >= 0 {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	nums1 := []int{-1, 3, 0, 0, 0, 0, 0}
	nums2 := []int{0, 0, 1, 2, 3}
	merge(nums1, 2, nums2, 5)
	fmt.Println(nums1)
}
