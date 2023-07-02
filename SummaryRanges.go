package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	// Declare answer variable
	ans := []string{}
	lenNums := len(nums)

	// handle input with length of 1 & 0
	if lenNums == 1 {
		strValue := strconv.Itoa(nums[0])
		ans = []string{strValue}
		return ans
	} else if lenNums == 0 {
		return ans
	}

	// Iterate through input
	// Variable i as the left pointer
	for i := 0; i < lenNums; {
		// Set right pointer to check the next variable
		j := i + 1
		// Check if the next number is a duplicate
		// If yes shift the left pointer to the right
		if nums[i] == nums[j] {
			i++
			//	Check if number in the left pointer and right pointer is a sequence
		} else if nums[j]-nums[i] == 1 {
			// Shift the right pointer to the right
			j++
			// Keep performing sequence checks
			for j != lenNums && nums[j]-nums[j-1] == 1 {
				j++
			}
			// Append the sequence to the answer array
			ans = append(ans, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
			// Shift the position of the left pointer to the current position of the right pointer
			i = j
			//	In case the number in left pointer and right pointer is not a duplicate or a sequence
		} else {
			// Insert the left pointer's value to the answer and shift it to the right
			ans = append(ans, strconv.Itoa(nums[i]))
			i++
		}

		// If the left pointer reach the end of the input,
		// Input the left pointer value to the answer
		if i == lenNums-1 {
			ans = append(ans, strconv.Itoa(nums[i]))
			i++
		}
	}
	return ans
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}
