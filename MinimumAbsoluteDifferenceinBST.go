package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt64
	prev := -1

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		if prev != -1 && node.Val-prev < minDiff {
			minDiff = node.Val - prev
		}
		prev = node.Val
		traverse(node.Right)
	}
	traverse(root)
	return minDiff
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}

	fmt.Println(getMinimumDifference(root))
}
