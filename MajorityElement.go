package main

import "fmt"

func majorityElement(nums []int) int {
	minLen := len(nums) / 2
	dict := map[int]int{}

	for _, num := range nums {
		dict[num] += 1
		if dict[num] > minLen {
			return num
		}
	}
	return 0
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
