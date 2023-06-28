package main

import "fmt"

func getAverages(nums []int, k int) []int {
	lenNums := len(nums)
	ArrRange := k*2 + 1

	// handle in case nums length is less than k
	if lenNums < ArrRange {
		for i, _ := range nums {
			nums[i] = -1
		}
		return nums
	}

	res := make([]int, lenNums)
	copy(res, nums)

	for i := k; i < lenNums-k; i++ {
		total := 0
		for j := i - k; j < i+k+1; j++ {
			total += nums[j]
		}

		res[i] = total / ArrRange
	}

	for i, j := 0, lenNums-1; i < k; i++ {
		res[i] = -1
		res[j] = -1
		j--
	}

	return res
}

func main() {
	nums := []int{8}
	k := 100000
	fmt.Println(getAverages(nums, k))
}
