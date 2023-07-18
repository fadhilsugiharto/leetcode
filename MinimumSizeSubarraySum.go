package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	lenN := len(nums)

	if lenN == 0 {
		return 0
	}

	lp, rp := 0, 0
	total := nums[lp]
	n := 1
	subArrLen := math.MaxInt64

	for lp <= lenN-1 {
		if total < target {
			rp++
			if rp > lenN-1 {
				break
			}
			n++
			total += nums[rp]
		} else {
			if n < subArrLen {
				subArrLen = n
			}
			total = total - nums[lp]
			lp++
			n--
		}
	}

	if subArrLen != math.MaxInt64 {
		return subArrLen
	}

	return 0
}

//func minSubArrayLen(target int, nums []int) int {
//	prefix := make([]int, len(nums)+1)
//
//	for i := 0; i < len(nums); i++ {
//		prefix[i+1] = prefix[i] + nums[i]
//	}
//
//	i := 0
//	j := 1
//
//	minlen := math.MaxInt
//
//	for j < len(prefix) {
//		diff := prefix[j] - prefix[i]
//		if diff < target {
//			j++
//		} else {
//			if j-i < minlen {
//				minlen = j - i
//			}
//			i++
//		}
//	}
//
//	if minlen == math.MaxInt {
//		return 0
//	}
//
//	return minlen
//}

func main() {
	nums := []int{10, 2, 3}
	fmt.Println(minSubArrayLen(6, nums))
}
