package main

import "fmt"

func strStr2(haystack string, needle string) int {
	nl := len(needle)
	if nl == 0 {
		return 0
	}
	hl := len(haystack)
	if hl == 0 {
		return -1
	}
	j := 0
	min := nl
	repeat := false
	for i := 1; i < nl; i++ {
		if needle[i] == needle[j] {
			j++
			repeat = true
		} else {
			if repeat {
				fmt.Println("end ", i, j)
				if min > j {
					min = j
				}
				j = 0
			}
		}
	}
	fmt.Println("min ", min, needle[:min])
	return 0

}
func main() {
	fmt.Println(strStr("abcdeabc", "abcab"))
	fmt.Println(strStr("abcdeabc", "abcabe"))
	fmt.Println(strStr("abcd e aebc", "abc"))
}
