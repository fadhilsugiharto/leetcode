package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	answer := [][]int{}
	i := 0
	dictAnswer := map[string]bool{}

	for i < len(nums) {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {

				// check if 3sum is equal to 0
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := []int{
						nums[i],
						nums[j],
						nums[k],
					}

					// sort list to help detect dups in ans
					sort.Ints(temp)

					// check if answer already exist
					key := fmt.Sprintf("%v", temp)
					if !dictAnswer[key] {
						answer = append(answer, temp)
						dictAnswer[key] = true
					}
				}
			}
		}

		i++
		if i == len(nums)-2 {
			return answer
		}
	}

	return answer
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
