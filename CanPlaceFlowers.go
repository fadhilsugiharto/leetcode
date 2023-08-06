package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	beds := len(flowerbed)
	i := 0

	for i < beds {
		if flowerbed[i] == 1 {
			i += 2
			continue
		}

		if i != 0 && flowerbed[i-1] == 1 {
			i += 1
			continue
		}

		if i+1 == beds || flowerbed[i+1] == 0 {
			n--
		}
		i += 2
	}

	if n <= 0 {
		return true
	}

	return false
}

func main() {
	flowerbed := []int{0, 1, 0, 1, 0, 1, 0, 0}
	fmt.Println(canPlaceFlowers(flowerbed, 1))
}
