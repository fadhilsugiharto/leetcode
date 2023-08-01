package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left := 1
	right := len(arr) - 2

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid-1] < arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return -1
}

func main() {
	arr := []int{3, 5, 3, 2, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}
