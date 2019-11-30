package main

import "fmt"

func convert(s string, numRows int) string {
	result := ""
	x := 0
	n := numRows
	for i, _ := range s {
		x = (i / (2*n - 2)) * (n - 1)
		x += i - x

		result += string(s[x])
	}
	return result
}

func main() {
	s := "LEETCODEISHIRING"
	fmt.Println(convert(s, 4))
}
