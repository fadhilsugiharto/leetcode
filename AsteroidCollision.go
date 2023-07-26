package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	remnant := []int{}

	for _, asteroid := range asteroids {
		if asteroid > 0 {
			remnant = append(remnant, asteroid)
		} else {
			if len(remnant) != 0 {

				ast := asteroid * -1
				for i := len(remnant) - 1; i >= 0; i-- {
					if remnant[len(remnant)-1] < 0 {
						remnant = append(remnant, asteroid)
						break
					}

					if remnant[i] < ast {
						remnant = remnant[:len(remnant)-1]
						if len(remnant) == 0 {
							remnant = append(remnant, asteroid)
						}
					} else if remnant[i] == ast {
						remnant = remnant[:len(remnant)-1]
						break
					} else {
						break
					}
				}
			} else {
				remnant = append(remnant, asteroid)
			}
		}
	}
	return remnant
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{-2, -2, 1, -2}))
}
