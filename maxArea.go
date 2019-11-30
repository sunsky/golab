package main

import "fmt"

func maxArea(height []int) int {
	l := len(height)
	max := [3]int{0, 0, 0}
	minY := 0
	for i, v := range height {
		for j := i + 1; j < l; j++ {
			if v > height[j] {
				minY = height[j]
			} else {
				minY = v
			}
			if a := (j - i) * minY; a > max[2] {
				max = [3]int{i, j, a}
			}
		}
	}
	fmt.Println("res ", max)
	return max[2]
}
func main() {
	fmt.Println("vim-go", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
}
