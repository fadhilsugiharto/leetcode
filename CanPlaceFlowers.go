package main

import "fmt"

//func canPlaceFlowers(flowerbed []int, n int) bool {
//	beds := len(flowerbed)
//	i := 0
//
//	for i < beds && n > 0 {
//		if flowerbed[i] == 1 {
//			i += 2
//			continue
//		}
//
//		if i != 0 && flowerbed[i-1] == 1 {
//			i += 1
//			continue
//		}
//
//		if i+1 == beds || flowerbed[i+1] == 0 {
//			n--
//		}
//		i += 2
//	}
//
//	return n <= 0
//}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			prevIsZero := i == 0 || flowerbed[i-1] == 0
			nextIsZero := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if prevIsZero && nextIsZero {
				flowerbed[i] = 1
				n--
				if n == 0 {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	fmt.Println(canPlaceFlowers(flowerbed, 2))
}
