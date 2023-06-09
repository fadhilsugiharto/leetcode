package main

import "fmt"

//func topKFrequent(nums []int, k int) []int {
//	dict := make(map[int]int)
//	ans := []int{}
//	keyToBeDeleted := 0
//
//	for _, num := range nums {
//		dict[num] += 1
//	}
//
//	for i := 0; i < k; i++ {
//		max := 0
//		for num, count := range dict {
//			if count > max {
//				max = count
//				keyToBeDeleted = num
//			}
//		}
//		ans = append(ans, keyToBeDeleted)
//		delete(dict, keyToBeDeleted)
//	}
//
//	return ans
//}

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int)
	ans := []int{}

	for _, num := range nums {
		dict[num]++
	}

	total := map[int][]int{}
	for num, val := range dict {
		total[val] = append(total[val], num)
	}

	for i := len(nums); len(ans) != k; i-- {
		for _, val := range total[i] {
			if len(ans) != k {
				ans = append(ans, val)
			}
		}
	}

	return ans
}

func main() {
	nums := []int{
		1, 1, 1, 2, 2, 3, 4, 4, 4,
	}
	fmt.Println(topKFrequent(nums, 1))
}
