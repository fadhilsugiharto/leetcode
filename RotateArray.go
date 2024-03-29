package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	rev(nums, 0, n-1)
	rev(nums, 0, k-1)
	rev(nums, k, n-1)
}

func rev(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

//func rotate(nums []int, k int) {
//	n := len(nums)
//	k %= n
//
//	for i := 0; i < k; i++ {
//		temp := nums[n-1]
//
//		for i := n - 1; i > 0; i-- {
//			nums[i] = nums[i-1]
//		}
//		nums[0] = temp
//	}
//
//}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
