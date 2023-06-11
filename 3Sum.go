package main

import (
	"fmt"
	"sort"
)

//func threeSum(nums []int) [][]int {
//	answer := [][]int{}
//	i := 0
//	dictAnswer := map[string]bool{}
//
//	for i < len(nums) {
//		for j := i + 1; j < len(nums); j++ {
//			for k := j + 1; k < len(nums); k++ {
//
//				// check if 3sum is equal to 0
//				if nums[i]+nums[j]+nums[k] == 0 {
//					temp := []int{
//						nums[i],
//						nums[j],
//						nums[k],
//					}
//
//					// sort list to help detect dups in ans
//					sort.Ints(temp)
//
//					// check if answer already exist
//					key := fmt.Sprintf("%v", temp)
//					if !dictAnswer[key] {
//						answer = append(answer, temp)
//						dictAnswer[key] = true
//					}
//				}
//			}
//		}
//
//		i++
//		if i == len(nums)-2 {
//			return answer
//		}
//	}
//
//	return answer
//}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	answer := [][]int{}

	i := 0
	for i < len(nums)-2 {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		target := -nums[i]
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				answer = append(answer, []int{-target, nums[left], nums[right]})
				left++
				right--
				for nums[left] == nums[left-1] && left < right {
					left++
				}
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
		i++
	}
	return answer
}

func main() {
	fmt.Println(threeSum([]int{1, -1, -1, 0}))
}
