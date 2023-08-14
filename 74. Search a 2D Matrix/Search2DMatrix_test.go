package main

import (
	"testing"
)

func testSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	res := SearchMatrix(matrix, target)
	t.Logf("test finished got result: %t", res)
}
