package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min int = -1

func calc(n *TreeNode, h int) int {
	fmt.Println(n, h, min)
	if n == nil {
		return h
	}

	h += 1
	l := calc(n.Left, h)
	if min == -1 {
		min = l
	}
	if n.Left != nil && min > l {
		min = l
	}
	r := calc(n.Right, h)
	if n.Right != nil && min > r {
		min = r
	}
	return min

}
func minDepth(root *TreeNode) int {
	h := 0
	return calc(root, h)
}
func main() {
	fmt.Println(minDepth(&TreeNode{0, nil, nil}))
}
