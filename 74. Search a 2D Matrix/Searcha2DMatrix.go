package main

import "fmt"

func SearchMatrix(matrix [][]int, target int) bool {
	leftOuter := 0
	rightOuter := len(matrix) - 1

	for leftOuter <= rightOuter {
		midOuter := leftOuter + (rightOuter-leftOuter)/2
		leftInner := 0
		rightInner := len(matrix[midOuter]) - 1

		if matrix[midOuter][leftInner] <= target && matrix[midOuter][rightInner] >= target {
			arr := matrix[midOuter]
			for leftInner <= rightInner {
				midInner := leftInner + (rightInner-leftInner)/2
				midNum := arr[midInner]

				if midNum == target {
					return true
				}

				if midNum > target {
					rightInner = midInner - 1
				} else {
					leftInner = midInner + 1
				}
			}
			return false
		}

		if matrix[midOuter][leftInner] > target {
			rightOuter = midOuter - 1
		} else {
			leftOuter = midOuter + 1
		}

	}

	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(SearchMatrix(matrix, 3))
}
