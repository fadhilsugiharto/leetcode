package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := math.MinInt
	sum := 0

	for _, num := range nums {
		sum += num
		if sum > res {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
