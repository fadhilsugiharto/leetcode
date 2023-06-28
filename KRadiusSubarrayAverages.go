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

	if k == 0 {
		return nums
	}

	res := make([]int, lenNums)
	copy(res, nums)

	total, lp, rp := 0, 0, 0
	for i := k; i < lenNums-k; i++ {
		if i != k {
			total = total - lp
			lp = nums[i-k]
			rp = nums[i+k]
			total = total + rp
		} else {
			for j := i - k; j < i+k+1; j++ {
				total += nums[j]
			}
			lp = nums[i-k]
			rp = nums[i+k]
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
	nums := []int{40527, 53696, 10730, 66491, 62141, 83909, 78635, 18560}
	k := 2
	fmt.Println(getAverages(nums, k))
}
