package main

import "fmt"

//func asteroidCollision(asteroids []int) []int {
//	remnant := []int{}
//
//	for _, asteroid := range asteroids {
//		if asteroid > 0 {
//			remnant = append(remnant, asteroid)
//		} else {
//			if len(remnant) != 0 {
//
//				ast := asteroid * -1
//				for i := len(remnant) - 1; i >= 0; i-- {
//					if remnant[len(remnant)-1] < 0 {
//						remnant = append(remnant, asteroid)
//						break
//					}
//
//					if remnant[i] < ast {
//						remnant = remnant[:len(remnant)-1]
//						if len(remnant) == 0 {
//							remnant = append(remnant, asteroid)
//						}
//					} else if remnant[i] == ast {
//						remnant = remnant[:len(remnant)-1]
//						break
//					} else {
//						break
//					}
//				}
//			} else {
//				remnant = append(remnant, asteroid)
//			}
//		}
//	}
//	return remnant
//}

//func asteroidCollision(asteroids []int) []int {
//	st, n := make([]int, 0), len(asteroids)
//	for i := 0; i < n; i++ {
//		same := false
//		for len(st) > 0 && asteroids[i] < 0 && asteroids[st[len(st)-1]] > 0 && asteroids[st[len(st)-1]]+asteroids[i] <= 0 {
//			sum := asteroids[st[len(st)-1]] + asteroids[i]
//			st = st[:len(st)-1]
//			if sum == 0 {
//				same = true
//				break
//			}
//		}
//		if same == false && (asteroids[i] > 0 || len(st) == 0 || asteroids[st[len(st)-1]] < 0) {
//			st = append(st, i)
//		}
//	}
//	for i := 0; i < len(st); i++ {
//		st[i] = asteroids[st[i]]
//	}
//	return st
//}

func asteroidCollision(asteroids []int) []int {
	remnants := []int{}

	for _, asteroid := range asteroids {
		collide := false
		n := len(remnants)
		incoming := asteroid * -1
		if n > 0 && remnants[n-1] > 0 && asteroid < 0 && remnants[n-1] <= incoming {
			for n > 0 {
				if remnants[n-1] < 0 {
					remnants = append(remnants, asteroid)
					collide = true
					break
				} else if remnants[n-1] == incoming {
					remnants = remnants[:len(remnants)-1]
					collide = true
					break
				} else if remnants[n-1] < incoming {
					remnants = remnants[:len(remnants)-1]
				} else {
					collide = true
					break
				}
				n--
			}
		}

		if collide == false && (len(remnants) == 0 || asteroid > 0 || remnants[len(remnants)-1] < 0) {
			remnants = append(remnants, asteroid)
		}
	}
	return remnants
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
}
