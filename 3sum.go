package main

import "fmt"

func cut(a int, arr []int) []int {
	others := arr
	for i, v := range arr {
		if a == v {
			//fmt.Println(a, others)
			if i+1 < len(arr) {
				others = append(others[:i], others[i+1:]...)
			} else {

				others = others[:i]
			}
		}
	}
	return others
}

func threeSum(ns []int) [][]int {
	const sum int = 0
	nsl := len(ns)
	tmp := [][]int{}
	for i, v := range ns {
		if i+1 >= nsl {
			break
		}
		for j, _v := range ns[i+1:] {
			if i+j+2 >= nsl {
				break
			}
			for _, __v := range ns[i+j+2:] {
				//fmt.Println(v, _v, __v)
				if v+_v+__v == sum {
					has := false
					t2 := []int{}
					for _, t := range tmp {
						t2 = []int{v, _v, __v}
						for _, tt := range t {
							t2 = cut(tt, t2)
						}
						if len(t2) == 0 {
							has = true
							break
						}
					}
					if !has {
						//sort.Ints(t2)
						tmp = append(tmp, []int{v, _v, __v})
					}
				}
			}
		}
	}
	return tmp

}
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-11, -3, -6, 12, -15, -13, -7, -3, 13, -2, -10, 3, 12, -12, 6, -6, 12, 9, -2, -12, 14, 11, -4, 11, -8, 8, 0, -12, 4, -5, 10, 8, 7, 11, -3, 7, 5, -3, -11, 3, 11, -13, 14, 8, 12, 5, -12, 10, -8, -7, 5, -9, -11, -14, 9, -12, 1, -6, -8, -10, 4, 9, 6, -3, -3, -12, 11, 9, 1, 8, -10, -3, 2, -11, -10, -1, 1, -15, -6, 8, -7, 6, 6, -10, 7, 0, -7, -7, 9, -8, -9, -9, -14, 12, -5, -10, -15, -9, -15, -7, 6, -10, 5, -7, -14, 3, 8, 2, 3, 9, -12, 4, 1, 9, 1, -15, -13, 9, -14, 11, 9}))
}
