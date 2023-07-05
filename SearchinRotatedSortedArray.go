package main

import "fmt"

func search(nums []int, target int) int {
	lenNums := len(nums)
	left := 0
	right := lenNums - 1

	for left >= 0 || right <= lenNums-1 {
		middle := (left + right) / 2
		if target == nums[middle] {
			return middle
		}
		if left >= right || right < left {
			break
		}
		// check if sequence is on the left
		if nums[left] < nums[middle] {
			if target >= nums[left] && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else /* if sequence is in the right */ {
			if target <= nums[right] && target >= nums[middle+1] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}
