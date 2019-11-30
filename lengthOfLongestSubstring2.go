package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	fmt.Println("input ", s)
	l := len(s)
	m := map[byte]int{}
	max := 0
	count := 0
	for i, j := 0, 0; i < l; i++ {
		if v, ok := m[s[i]]; ok && v > j-1 {
			j = v + 1
		}
		count = i - j + 1
		if max < count {
			max = count
		}
		m[s[i]] = i
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
