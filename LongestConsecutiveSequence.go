package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	dict := map[int]bool{}
	lenNums := len(nums)

	if lenNums == 0 {
		return 0
	}

	for _, num := range nums {
		dict[num] = true
	}

	sum := 1

	for _, num := range nums {
		length := 1
		if _, ok := dict[num-1]; !ok {
			for {
				if _, y := dict[num+1]; y {
					length++
					if length > sum {
						sum = length
					}
					num++
				} else {
					break
				}
			}
		}
	}

	return sum
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
