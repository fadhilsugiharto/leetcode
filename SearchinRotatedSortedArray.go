package main

import "fmt"

func search(nums []int, target int) int {
	lenNums := len(nums)
	left := 0
	right := lenNums - 1

	for left <= right {
		middle := left + (right-left)/2
		if target == nums[middle] {
			return middle
		}

		// check if sequence is on the right
		if nums[left] > nums[middle] {
			if nums[middle] < target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else /* if sequence is on the left */ {
			if nums[left] <= target && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}
