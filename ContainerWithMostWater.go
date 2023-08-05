package main

import "fmt"

func min(left int, right int) int {
	if left < right {
		return left
	}
	return right
}

func maxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1
	res := 0

	for left < right {
		leftNum := height[left]
		rightNum := height[right]

		total := min(leftNum, rightNum) * (right - left)

		if total > res {
			res = total
		}

		if height[left] < height[right] {
			for left < n {
				if height[left] <= leftNum {
					left++
				} else {
					break
				}
			}
		} else {
			for right >= 0 {
				if height[right] <= rightNum {
					right--
				} else {
					break
				}
			}
		}
	}

	return res
}

func main() {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	fmt.Println(maxArea(height))
}
