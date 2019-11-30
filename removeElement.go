package main

import "fmt"

func removeElement(nums []int, val int) int {
	s := len(nums)
	if s == 0 {
		return 0
	}
	i := 0
	for j := 0; j < s; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	fmt.Println(nums)
	return i

}
func main() {
	//, val = 3,
	nums := []int{3, 2, 2, 3}
	fmt.Println(2, removeElement(nums, 3))
}
