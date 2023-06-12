package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	for i := 0; i < len(coordinates)-2; i++ {
		x0, x1, x2 := coordinates[i][0], coordinates[i+1][0], coordinates[i+2][0]
		y0, y1, y2 := coordinates[i][1], coordinates[i+1][1], coordinates[i+2][1]
		if (x1-x0)*(y2-y1) != (y1-y0)*(x2-x1) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkStraightLine([][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7},
	}))
}
