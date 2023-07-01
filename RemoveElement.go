package main

import "fmt"

//func removeElements(nums []int, val int) int {
//	for i, num := range nums {
//		if num == val {
//			j := i + 1
//			for j < len(nums) && nums[j] == val {
//				j++
//			}
//			if j == len(nums) {
//				nums = nums[:i]
//			} else {
//				nums = append(nums[:i], nums[j:]...)
//			}
//		}
//		fmt.Println(nums)
//	}
//	return len(nums)
//}

func removeElement(nums []int, val int) int {
	count := 0
	for _, num := range nums {
		if num != val {
			nums[count] = num
			count++
		}
	}
	return count
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}
