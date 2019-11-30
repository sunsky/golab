package main

import "fmt"

func buble(a []int) []int {
	l := len(a)
	fmt.Println("inputed ", a, " len ", l)
	pos := 0
	last := l - 1
	for i := l; i >= 0; i-- {
		swap := 0
		for j := 0; j < last; j++ {
			if a[j] > a[j+1] {
				swap = 1
				pos = j
				a[j] ^= a[j+1]
				a[j+1] ^= a[j]
				a[j] ^= a[j+1]
			}
		}
		last = pos
		if swap == 0 {
			return a
		}
	}
	return a
}
func main() {
	a := []int{3, 7, 8, 11, 99}
	fmt.Println("len ", len(a), buble(a))
	a = []int{3, 7, 1, 4, 8, 2, 5, 11, 99, 7}
	fmt.Println("len ", len(a), buble(a))
}
