package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	s := len(nums)
	l := 0
	for r := 1; r < s; r++ {
		if nums[l] != nums[r] {
			//fmt.Println(l, r)
			l++
			if l != r {
				nums[l] = nums[r]
			}
		}
	}
	l++
	//fmt.Println(nums)
	return l
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(5, removeDuplicates(nums))
}
