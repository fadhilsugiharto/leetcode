package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	ans := []int{}
	keyToBeDeleted := 0

	for _, num := range nums {
		dict[num] += 1
	}

	for i := 0; i < k; i++ {
		max := 0
		for num, count := range dict {
			if count > max {
				max = count
				keyToBeDeleted = num
			}
		}
		ans = append(ans, keyToBeDeleted)
		delete(dict, keyToBeDeleted)
	}

	return ans
}

func main() {
	nums := []int{
		3, 0, 1, 0,
	}
	fmt.Println(topKFrequent(nums, 1))
}
