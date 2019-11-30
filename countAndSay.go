package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	step := "1"
	for i := 1; i < n; i++ {
		sl := len(step)
		start := 0
		next := ""
		for j := 0; j < sl; j++ {
			if step[start] != step[j] {
				next += strconv.Itoa(j-start) + string(step[start])
				start = j
			}
			if j == sl-1 {
				next += strconv.Itoa(j-start+1) + string(step[start])
			}

		}

		step = next
	}
	return step
}
func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))
	fmt.Println(countAndSay(8))
	fmt.Println(countAndSay(9))
	fmt.Println(countAndSay(10))
}
