package main

import "fmt"

func maxProfit(prices []int) int {
	profit, left, right := 0, 0, 1

	for right < len(prices) {
		diff := prices[right] - prices[left]
		if prices[right] > prices[left] {
			if profit < diff {
				profit = diff
			}
		} else {
			left = right
		}
		right++
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 2, 3, 1, 5, 3, 6, 4}))
}
