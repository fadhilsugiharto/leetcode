package main

import "fmt"

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)
	for _, num := range nums {
		if dict[num] {
			return true
		}
		dict[num] = true
	}
	return false
}

func main() {
	nums := []int{
		1, 5, -2, -4, 0, -2,
	}
	fmt.Println(containsDuplicate(nums))
}
