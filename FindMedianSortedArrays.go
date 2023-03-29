package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var median float64
	var combined []float64
	combinedInt := append(nums1, nums2...)
	sort.Slice(combinedInt, func(i, j int) bool {
		return combinedInt[i] < combinedInt[j]
	})
	lenCombined := len(combinedInt)

	for _, num := range combinedInt {
		combined = append(combined, float64(num))
	}

	if lenCombined%2 == 0 {
		median = (combined[int(lenCombined/2)-1] + combined[lenCombined/2]) / 2
	} else {
		median = combined[lenCombined/2]
	}

	return median
}

func main() {
	nums1 := []int{1, 7}
	nums2 := []int{3, 9}

	median := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("%.5f", median)
}
